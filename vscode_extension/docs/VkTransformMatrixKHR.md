# VkTransformMatrixKHR(3) Manual Page

## Name

VkTransformMatrixKHR - Structure specifying a 3x4 affine transformation
matrix



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkTransformMatrixKHR` structure is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkTransformMatrixKHR {
    float    matrix[3][4];
} VkTransformMatrixKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkTransformMatrixKHR VkTransformMatrixNV;
```

## <a href="#_members" class="anchor"></a>Members

- `matrix` is a 3x4 row-major affine transformation matrix.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkTransformMatrixKHR-matrix-03799"
  id="VUID-VkTransformMatrixKHR-matrix-03799"></a>
  VUID-VkTransformMatrixKHR-matrix-03799  
  The first three columns of `matrix` **must** define an invertible 3x3
  matrix

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html),
[VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTransformMatrixKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
