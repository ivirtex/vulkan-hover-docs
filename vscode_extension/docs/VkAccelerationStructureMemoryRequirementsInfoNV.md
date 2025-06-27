# VkAccelerationStructureMemoryRequirementsInfoNV(3) Manual Page

## Name

VkAccelerationStructureMemoryRequirementsInfoNV - Structure specifying
acceleration to query for memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureMemoryRequirementsInfoNV` structure is
defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkAccelerationStructureMemoryRequirementsInfoNV {
    VkStructureType                                    sType;
    const void*                                        pNext;
    VkAccelerationStructureMemoryRequirementsTypeNV    type;
    VkAccelerationStructureNV                          accelerationStructure;
} VkAccelerationStructureMemoryRequirementsInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` selects the type of memory requirement being queried.
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV` returns
  the memory requirements for the object itself.
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV`
  returns the memory requirements for the scratch memory when doing a
  build.
  `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV`
  returns the memory requirements for the scratch memory when doing an
  update.

- `accelerationStructure` is the acceleration structure to be queried
  for memory requirements.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-sType-sType"
  id="VUID-VkAccelerationStructureMemoryRequirementsInfoNV-sType-sType"></a>
  VUID-VkAccelerationStructureMemoryRequirementsInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_INFO_NV`

- <a
  href="#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-pNext-pNext"
  id="VUID-VkAccelerationStructureMemoryRequirementsInfoNV-pNext-pNext"></a>
  VUID-VkAccelerationStructureMemoryRequirementsInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-type-parameter"
  id="VUID-VkAccelerationStructureMemoryRequirementsInfoNV-type-parameter"></a>
  VUID-VkAccelerationStructureMemoryRequirementsInfoNV-type-parameter  
  `type` **must** be a valid
  [VkAccelerationStructureMemoryRequirementsTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html)
  value

- <a
  href="#VUID-VkAccelerationStructureMemoryRequirementsInfoNV-accelerationStructure-parameter"
  id="VUID-VkAccelerationStructureMemoryRequirementsInfoNV-accelerationStructure-parameter"></a>
  VUID-VkAccelerationStructureMemoryRequirementsInfoNV-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureMemoryRequirementsTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsTypeNV.html),
[VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMemoryRequirementsInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
