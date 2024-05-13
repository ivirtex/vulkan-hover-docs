# VkAccelerationStructureMemoryRequirementsTypeNV(3) Manual Page

## Name

VkAccelerationStructureMemoryRequirementsTypeNV - Acceleration structure
memory requirement type



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of `type` in
`VkAccelerationStructureMemoryRequirementsInfoNV` are:,

``` c
// Provided by VK_NV_ray_tracing
typedef enum VkAccelerationStructureMemoryRequirementsTypeNV {
    VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV = 0,
    VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV = 1,
    VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV = 2,
} VkAccelerationStructureMemoryRequirementsTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_OBJECT_NV`
  requests the memory requirement for the `VkAccelerationStructureNV`
  backing store.

- `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV`
  requests the memory requirement for scratch space during the initial
  build.

- `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV`
  requests the memory requirement for scratch space during an update.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMemoryRequirementsTypeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
