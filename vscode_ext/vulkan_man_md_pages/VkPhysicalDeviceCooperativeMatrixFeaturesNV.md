# VkPhysicalDeviceCooperativeMatrixFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceCooperativeMatrixFeaturesNV - Structure describing
cooperative matrix features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCooperativeMatrixFeaturesNV` structure is defined
as:

``` c
// Provided by VK_NV_cooperative_matrix
typedef struct VkPhysicalDeviceCooperativeMatrixFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           cooperativeMatrix;
    VkBool32           cooperativeMatrixRobustBufferAccess;
} VkPhysicalDeviceCooperativeMatrixFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-cooperativeMatrix-NV"></span> `cooperativeMatrix`
  indicates that the implementation supports the `CooperativeMatrixNV`
  SPIR-V capability.

- <span id="features-cooperativeMatrixRobustBufferAccess-NV"></span>
  `cooperativeMatrixRobustBufferAccess` indicates that the
  implementation supports robust buffer access for SPIR-V
  `OpCooperativeMatrixLoadNV` and `OpCooperativeMatrixStoreNV`
  instructions.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceCooperativeMatrixFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceCooperativeMatrixFeaturesNV` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceCooperativeMatrixFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceCooperativeMatrixFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceCooperativeMatrixFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cooperative_matrix.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCooperativeMatrixFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
