# VkCooperativeVectorMatrixLayoutNV(3) Manual Page

## Name

VkCooperativeVectorMatrixLayoutNV - Specify cooperative vector matrix layout



## [](#_c_specification)C Specification

Possible values for [VkCooperativeVectorMatrixLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeVectorMatrixLayoutNV.html) include:

```c++
// Provided by VK_NV_cooperative_vector
typedef enum VkCooperativeVectorMatrixLayoutNV {
    VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_ROW_MAJOR_NV = 0,
    VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_COLUMN_MAJOR_NV = 1,
    VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_INFERENCING_OPTIMAL_NV = 2,
    VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_TRAINING_OPTIMAL_NV = 3,
} VkCooperativeVectorMatrixLayoutNV;
```

## [](#_description)Description

- `VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_ROW_MAJOR_NV` corresponds to SPIR-V `RowMajorNV` layout.
- `VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_COLUMN_MAJOR_NV` corresponds to SPIR-V `ColumnMajorNV` layout.
- `VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_INFERENCING_OPTIMAL_NV` corresponds to SPIR-V `InferencingOptimalNV` layout.
- `VK_COOPERATIVE_VECTOR_MATRIX_LAYOUT_TRAINING_OPTIMAL_NV` corresponds to SPIR-V `TrainingOptimalNV` layout.

All enum values match the corresponding SPIR-V value.

Row-major layout has elements of each row stored consecutively in memory, with a controllable stride from the start of one row to the start of the next row. Column-major layout has elements of each column stored consecutively in memory, with a controllable stride from the start of one column to the start of the next column. Inferencing-optimal and Training-optimal layouts are implementation-dependent, and the application **can** convert a matrix to those layouts using [vkConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkConvertCooperativeVectorMatrixNV.html) or [vkCmdConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdConvertCooperativeVectorMatrixNV.html). Training-optimal layout with `VK_COMPONENT_TYPE_FLOAT16_KHR` or `VK_COMPONENT_TYPE_FLOAT32_KHR` type has the additional guarantee that the application **can** reinterpret the data as an array of elements and perform element-wise operations on the data, and finite values in any padding elements do not affect the result of a matrix-vector multiply (inf/NaN values **may** still cause NaN values in the result).

## [](#_see_also)See Also

[VK\_NV\_cooperative\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_vector.html), [VkConvertCooperativeVectorMatrixInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConvertCooperativeVectorMatrixInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCooperativeVectorMatrixLayoutNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0