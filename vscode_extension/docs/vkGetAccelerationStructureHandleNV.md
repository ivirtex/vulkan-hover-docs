# vkGetAccelerationStructureHandleNV(3) Manual Page

## Name

vkGetAccelerationStructureHandleNV - Get opaque acceleration structure handle



## [](#_c_specification)C Specification

To allow constructing geometry instances with device code if desired, we need to be able to query an opaque handle for an acceleration structure. This handle is a value of 8 bytes. To get this handle, call:

```c++
// Provided by VK_NV_ray_tracing
VkResult vkGetAccelerationStructureHandleNV(
    VkDevice                                    device,
    VkAccelerationStructureNV                   accelerationStructure,
    size_t                                      dataSize,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the acceleration structures.
- `accelerationStructure` is the acceleration structure.
- `dataSize` is the size in bytes of the buffer pointed to by `pData`.
- `pData` is a pointer to an application-allocated buffer where the results will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetAccelerationStructureHandleNV-dataSize-02240)VUID-vkGetAccelerationStructureHandleNV-dataSize-02240  
  `dataSize` **must** be large enough to contain the result of the query, as described above
- [](#VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-02787)VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-02787  
  `accelerationStructure` **must** be bound completely and contiguously to a single `VkDeviceMemory` object via [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html)

Valid Usage (Implicit)

- [](#VUID-vkGetAccelerationStructureHandleNV-device-parameter)VUID-vkGetAccelerationStructureHandleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parameter)VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-vkGetAccelerationStructureHandleNV-pData-parameter)VUID-vkGetAccelerationStructureHandleNV-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkGetAccelerationStructureHandleNV-dataSize-arraylength)VUID-vkGetAccelerationStructureHandleNV-dataSize-arraylength  
  `dataSize` **must** be greater than `0`
- [](#VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parent)VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parent  
  `accelerationStructure` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetAccelerationStructureHandleNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0