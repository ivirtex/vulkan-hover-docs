# vkBindAccelerationStructureMemoryNV(3) Manual Page

## Name

vkBindAccelerationStructureMemoryNV - Bind acceleration structure memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To attach memory to one or more acceleration structures at a time, call:

``` c
// Provided by VK_NV_ray_tracing
VkResult vkBindAccelerationStructureMemoryNV(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindAccelerationStructureMemoryInfoNV* pBindInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the acceleration structures
  and memory.

- `bindInfoCount` is the number of elements in `pBindInfos`.

- `pBindInfos` is a pointer to an array of
  [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindAccelerationStructureMemoryInfoNV.html)
  structures describing acceleration structures and memory to bind.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkBindAccelerationStructureMemoryNV-device-parameter"
  id="VUID-vkBindAccelerationStructureMemoryNV-device-parameter"></a>
  VUID-vkBindAccelerationStructureMemoryNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindAccelerationStructureMemoryNV-pBindInfos-parameter"
  id="VUID-vkBindAccelerationStructureMemoryNV-pBindInfos-parameter"></a>
  VUID-vkBindAccelerationStructureMemoryNV-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of
  `bindInfoCount` valid
  [VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindAccelerationStructureMemoryInfoNV.html)
  structures

- <a
  href="#VUID-vkBindAccelerationStructureMemoryNV-bindInfoCount-arraylength"
  id="VUID-vkBindAccelerationStructureMemoryNV-bindInfoCount-arraylength"></a>
  VUID-vkBindAccelerationStructureMemoryNV-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkBindAccelerationStructureMemoryInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindAccelerationStructureMemoryInfoNV.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindAccelerationStructureMemoryNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
