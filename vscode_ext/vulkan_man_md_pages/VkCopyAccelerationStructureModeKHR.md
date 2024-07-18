# VkCopyAccelerationStructureModeKHR(3) Manual Page

## Name

VkCopyAccelerationStructureModeKHR - Acceleration structure copy mode



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of `mode` specifying additional operations to perform
during the copy, are:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkCopyAccelerationStructureModeKHR {
    VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR = 0,
    VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR = 1,
    VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR = 2,
    VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR = 3,
  // Provided by VK_NV_ray_tracing
    VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV = VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR,
  // Provided by VK_NV_ray_tracing
    VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV = VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR,
} VkCopyAccelerationStructureModeKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkCopyAccelerationStructureModeKHR VkCopyAccelerationStructureModeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR` creates a direct copy
  of the acceleration structure specified in `src` into the one
  specified by `dst`. The `dst` acceleration structure **must** have
  been created with the same parameters as `src`. If `src` contains
  references to other acceleration structures, `dst` will reference the
  same acceleration structures.

- `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR` creates a more
  compact version of an acceleration structure `src` into `dst`. The
  acceleration structure `dst` **must** have been created with a size at
  least as large as that returned by
  [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html)
  or
  [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteAccelerationStructuresPropertiesKHR.html)
  after the build of the acceleration structure specified by `src`. If
  `src` contains references to other acceleration structures, `dst` will
  reference the same acceleration structures.

- `VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR` serializes the
  acceleration structure to a semi-opaque format which can be reloaded
  on a compatible implementation.

- `VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR` deserializes the
  semi-opaque serialization format in the buffer to the acceleration
  structure.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureInfoKHR.html),
[VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html),
[VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html),
[vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyAccelerationStructureModeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
