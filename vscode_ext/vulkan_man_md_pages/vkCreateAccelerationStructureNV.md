# vkCreateAccelerationStructureNV(3) Manual Page

## Name

vkCreateAccelerationStructureNV - Create a new acceleration structure
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create acceleration structures, call:

``` c
// Provided by VK_NV_ray_tracing
VkResult vkCreateAccelerationStructureNV(
    VkDevice                                    device,
    const VkAccelerationStructureCreateInfoNV*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkAccelerationStructureNV*                  pAccelerationStructure);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the buffer object.

- `pCreateInfo` is a pointer to a
  [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html)
  structure containing parameters affecting creation of the acceleration
  structure.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pAccelerationStructure` is a pointer to a
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle in
  which the resulting acceleration structure object is returned.

## <a href="#_description" class="anchor"></a>Description

Similarly to other objects in Vulkan, the acceleration structure
creation merely creates an object with a specific “shape” as specified
by the information in
[VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html) and
`compactedSize` in `pCreateInfo`.

Once memory has been bound to the acceleration structure using
[vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindAccelerationStructureMemoryNV.html),
that memory is populated by calls to
[vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructureNV.html)
and
[vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureNV.html).

Acceleration structure creation uses the count and type information from
the geometries, but does not use the data references in the structures.

Valid Usage (Implicit)

- <a href="#VUID-vkCreateAccelerationStructureNV-device-parameter"
  id="VUID-vkCreateAccelerationStructureNV-device-parameter"></a>
  VUID-vkCreateAccelerationStructureNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateAccelerationStructureNV-pCreateInfo-parameter"
  id="VUID-vkCreateAccelerationStructureNV-pCreateInfo-parameter"></a>
  VUID-vkCreateAccelerationStructureNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html)
  structure

- <a href="#VUID-vkCreateAccelerationStructureNV-pAllocator-parameter"
  id="VUID-vkCreateAccelerationStructureNV-pAllocator-parameter"></a>
  VUID-vkCreateAccelerationStructureNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateAccelerationStructureNV-pAccelerationStructure-parameter"
  id="VUID-vkCreateAccelerationStructureNV-pAccelerationStructure-parameter"></a>
  VUID-vkCreateAccelerationStructureNV-pAccelerationStructure-parameter  
  `pAccelerationStructure` **must** be a valid pointer to a
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateAccelerationStructureNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
