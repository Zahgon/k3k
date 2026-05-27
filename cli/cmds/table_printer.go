package cmds

import (
	"k8s.io/apimachinery/pkg/runtime"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// createTable creates a table to print from the printerColumn defined in the CRD spec, plus the name at the beginning
func createTable[T runtime.Object](crd *apiextensionsv1.CustomResourceDefinition, objs []T) *metav1.Table {
	_ = "STUB: not implemented"
	return nil
}

func getPrinterColumnsFromCRD(crd *apiextensionsv1.CustomResourceDefinition) []apiextensionsv1.CustomResourceColumnDefinition {
	_ = "STUB: not implemented"
	return nil
}

func convertToTableColumns(printerColumns []apiextensionsv1.CustomResourceColumnDefinition) []metav1.TableColumnDefinition {
	_ = "STUB: not implemented"
	return nil
}

func createTableRows[T runtime.Object](objs []T, printerColumns []apiextensionsv1.CustomResourceColumnDefinition) []metav1.TableRow {
	_ = "STUB: not implemented"
	return nil
}

func buildRowCells(objMap map[string]any, printerColumns []apiextensionsv1.CustomResourceColumnDefinition) []any {
	_ = "STUB: not implemented"
	return nil
}

func toPointerSlice[T any](v []T) []*T { _ = "STUB: not implemented"; return nil }
