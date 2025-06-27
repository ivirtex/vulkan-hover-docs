# vkGetAccelerationStructureHandleNV(3) Manual Page

## Name

vkGetAccelerationStructureHandleNV - Get opaque acceleration structure
handle



## <a href="#_c_specification" class="anchor"></a>C Specification

To allow constructing geometry instances with device code if desired, we
need to be able to query an opaque handle for an acceleration structure.
This handle is a value of 8 bytes. To get this handle, call:

``` c
// Provided by VK_NV_ray_tracing
VkResult vkGetAccelerationStructureHandleNV(
    VkDevice                                    device,
    VkAccelerationStructureNV                   accelerationStructure,
    size_t                                      dataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the acceleration structures.

- `accelerationStructure` is the acceleration structure.

- `dataSize` is the size in bytes of the buffer pointed to by `pData`.

- `pData` is a pointer to an application-allocated buffer where the
  results will be written.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetAccelerationStructureHandleNV-dataSize-02240"
  id="VUID-vkGetAccelerationStructureHandleNV-dataSize-02240"></a>
  VUID-vkGetAccelerationStructureHandleNV-dataSize-02240  
  `dataSize` **must** be large enough to contain the result of the
  query, as described above

- <a
  href="#VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-02787"
  id="VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-02787"></a>
  VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-02787  
  `accelerationStructure` **must** be bound completely and contiguously
  to a single `VkDeviceMemory` object via
  [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html)

Valid Usage (Implicit)

- <a href="#VUID-vkGetAccelerationStructureHandleNV-device-parameter"
  id="VUID-vkGetAccelerationStructureHandleNV-device-parameter"></a>
  VUID-vkGetAccelerationStructureHandleNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parameter"
  id="VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parameter"></a>
  VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

- <a href="#VUID-vkGetAccelerationStructureHandleNV-pData-parameter"
  id="VUID-vkGetAccelerationStructureHandleNV-pData-parameter"></a>
  VUID-vkGetAccelerationStructureHandleNV-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a href="#VUID-vkGetAccelerationStructureHandleNV-dataSize-arraylength"
  id="VUID-vkGetAccelerationStructureHandleNV-dataSize-arraylength"></a>
  VUID-vkGetAccelerationStructureHandleNV-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a
  href="#VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parent"
  id="VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parent"></a>
  VUID-vkGetAccelerationStructureHandleNV-accelerationStructure-parent  
  `accelerationStructure` **must** have been created, allocated, or
  retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetAccelerationStructureHandleNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
