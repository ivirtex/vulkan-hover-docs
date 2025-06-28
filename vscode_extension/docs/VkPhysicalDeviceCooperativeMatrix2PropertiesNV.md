# VkPhysicalDeviceCooperativeMatrix2PropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceCooperativeMatrix2PropertiesNV - Structure describing cooperative matrix properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCooperativeMatrix2PropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_cooperative_matrix2
typedef struct VkPhysicalDeviceCooperativeMatrix2PropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           cooperativeMatrixWorkgroupScopeMaxWorkgroupSize;
    uint32_t           cooperativeMatrixFlexibleDimensionsMaxDimension;
    uint32_t           cooperativeMatrixWorkgroupScopeReservedSharedMemory;
} VkPhysicalDeviceCooperativeMatrix2PropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`cooperativeMatrixWorkgroupScopeMaxWorkgroupSize` is the maximum number of invocations in a workgroup when the module uses `OpTypeCooperativeMatrixKHR` with `Scope` equal to `Workgroup`.
- []()`cooperativeMatrixFlexibleDimensionsMaxDimension` is the maximum supported dimension for cooperative matrix types when the [`cooperativeMatrixFlexibleDimensions`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-cooperativeMatrixFlexibleDimensions) feature is enabled.
- []()`cooperativeMatrixWorkgroupScopeReservedSharedMemory` is the number of bytes of shared memory reserved for the implementation when the module uses `OpTypeCooperativeMatrixKHR` with `Scope` equal to `Workgroup`.

## [](#_description)Description

If the `VkPhysicalDeviceCooperativeMatrix2PropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCooperativeMatrix2PropertiesNV-sType-sType)VUID-VkPhysicalDeviceCooperativeMatrix2PropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_2_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_cooperative\_matrix2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_matrix2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCooperativeMatrix2PropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0