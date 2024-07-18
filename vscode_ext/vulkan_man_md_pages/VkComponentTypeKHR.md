# VkComponentTypeKHR(3) Manual Page

## Name

VkComponentTypeKHR - Specify SPIR-V cooperative matrix component type



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values for [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeKHR.html)
include:

``` c
// Provided by VK_KHR_cooperative_matrix
typedef enum VkComponentTypeKHR {
    VK_COMPONENT_TYPE_FLOAT16_KHR = 0,
    VK_COMPONENT_TYPE_FLOAT32_KHR = 1,
    VK_COMPONENT_TYPE_FLOAT64_KHR = 2,
    VK_COMPONENT_TYPE_SINT8_KHR = 3,
    VK_COMPONENT_TYPE_SINT16_KHR = 4,
    VK_COMPONENT_TYPE_SINT32_KHR = 5,
    VK_COMPONENT_TYPE_SINT64_KHR = 6,
    VK_COMPONENT_TYPE_UINT8_KHR = 7,
    VK_COMPONENT_TYPE_UINT16_KHR = 8,
    VK_COMPONENT_TYPE_UINT32_KHR = 9,
    VK_COMPONENT_TYPE_UINT64_KHR = 10,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_FLOAT16_NV = VK_COMPONENT_TYPE_FLOAT16_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_FLOAT32_NV = VK_COMPONENT_TYPE_FLOAT32_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_FLOAT64_NV = VK_COMPONENT_TYPE_FLOAT64_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_SINT8_NV = VK_COMPONENT_TYPE_SINT8_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_SINT16_NV = VK_COMPONENT_TYPE_SINT16_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_SINT32_NV = VK_COMPONENT_TYPE_SINT32_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_SINT64_NV = VK_COMPONENT_TYPE_SINT64_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_UINT8_NV = VK_COMPONENT_TYPE_UINT8_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_UINT16_NV = VK_COMPONENT_TYPE_UINT16_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_UINT32_NV = VK_COMPONENT_TYPE_UINT32_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_COMPONENT_TYPE_UINT64_NV = VK_COMPONENT_TYPE_UINT64_KHR,
} VkComponentTypeKHR;
```

or the equivalent

``` c
// Provided by VK_NV_cooperative_matrix
typedef VkComponentTypeKHR VkComponentTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COMPONENT_TYPE_FLOAT16_KHR` corresponds to SPIR-V `OpTypeFloat`
  16.

- `VK_COMPONENT_TYPE_FLOAT32_KHR` corresponds to SPIR-V `OpTypeFloat`
  32.

- `VK_COMPONENT_TYPE_FLOAT64_KHR` corresponds to SPIR-V `OpTypeFloat`
  64.

- `VK_COMPONENT_TYPE_SINT8_KHR` corresponds to SPIR-V `OpTypeInt` 8 0/1.

- `VK_COMPONENT_TYPE_SINT16_KHR` corresponds to SPIR-V `OpTypeInt` 16
  0/1.

- `VK_COMPONENT_TYPE_SINT32_KHR` corresponds to SPIR-V `OpTypeInt` 32
  0/1.

- `VK_COMPONENT_TYPE_SINT64_KHR` corresponds to SPIR-V `OpTypeInt` 64
  0/1.

- `VK_COMPONENT_TYPE_UINT8_KHR` corresponds to SPIR-V `OpTypeInt` 8 0/1.

- `VK_COMPONENT_TYPE_UINT16_KHR` corresponds to SPIR-V `OpTypeInt` 16
  0/1.

- `VK_COMPONENT_TYPE_UINT32_KHR` corresponds to SPIR-V `OpTypeInt` 32
  0/1.

- `VK_COMPONENT_TYPE_UINT64_KHR` corresponds to SPIR-V `OpTypeInt` 64
  0/1.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_cooperative_matrix.html),
[VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkComponentTypeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
