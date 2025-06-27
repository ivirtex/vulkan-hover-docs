# vkGetAccelerationStructureMemoryRequirementsNV(3) Manual Page

## Name

vkGetAccelerationStructureMemoryRequirementsNV - Get acceleration
structure memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

An acceleration structure has memory requirements for the structure
object itself, scratch space for the build, and scratch space for the
update.

Scratch space is allocated as a `VkBuffer`, so for
`VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV`
and
`VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV`
the `pMemoryRequirements->alignment` and
`pMemoryRequirements->memoryTypeBits` values returned by this call
**must** be filled with zero, and **should** be ignored by the
application.

To query the memory requirements, call:

``` c
// Provided by VK_NV_ray_tracing
void vkGetAccelerationStructureMemoryRequirementsNV(
    VkDevice                                    device,
    const VkAccelerationStructureMemoryRequirementsInfoNV* pInfo,
    VkMemoryRequirements2KHR*                   pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device on which the acceleration structure was
  created.

- `pInfo` is a pointer to a
  [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)
  structure specifying the acceleration structure to get memory
  requirements for.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2KHR.html) structure in
  which the requested acceleration structure memory requirements are
  returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetAccelerationStructureMemoryRequirementsNV-device-parameter"
  id="VUID-vkGetAccelerationStructureMemoryRequirementsNV-device-parameter"></a>
  VUID-vkGetAccelerationStructureMemoryRequirementsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetAccelerationStructureMemoryRequirementsNV-pInfo-parameter"
  id="VUID-vkGetAccelerationStructureMemoryRequirementsNV-pInfo-parameter"></a>
  VUID-vkGetAccelerationStructureMemoryRequirementsNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)
  structure

- <a
  href="#VUID-vkGetAccelerationStructureMemoryRequirementsNV-pMemoryRequirements-parameter"
  id="VUID-vkGetAccelerationStructureMemoryRequirementsNV-pMemoryRequirements-parameter"></a>
  VUID-vkGetAccelerationStructureMemoryRequirementsNV-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2KHR.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetAccelerationStructureMemoryRequirementsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
