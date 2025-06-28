# VkCopyAccelerationStructureModeKHR(3) Manual Page

## Name

VkCopyAccelerationStructureModeKHR - Acceleration structure copy mode



## [](#_c_specification)C Specification

Possible values of `mode` specifying additional operations to perform during the copy, are:

```c++
// Provided by VK_KHR_acceleration_structure
typedef enum VkCopyAccelerationStructureModeKHR {
    VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR = 0,
    VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR = 1,
  // Provided by VK_KHR_acceleration_structure
    VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR = 2,
  // Provided by VK_KHR_acceleration_structure
    VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR = 3,
  // Provided by VK_NV_ray_tracing
    VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV = VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR,
  // Provided by VK_NV_ray_tracing
    VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV = VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR,
} VkCopyAccelerationStructureModeKHR;
```

or the equivalent

```c++
// Provided by VK_NV_ray_tracing
typedef VkCopyAccelerationStructureModeKHR VkCopyAccelerationStructureModeNV;
```

## [](#_description)Description

- `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR` creates a direct copy of the acceleration structure specified in `src` into the one specified by `dst`. The `dst` acceleration structure **must** have been created with the same parameters as `src`. If `src` contains references to other acceleration structures, `dst` will reference the same acceleration structures.
- `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR` creates a more compact version of an acceleration structure `src` into `dst`. The acceleration structure `dst` **must** have been created with a size at least as large as that returned by [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html) , [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html), or [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteAccelerationStructuresPropertiesKHR.html) after the build of the acceleration structure specified by `src`. If `src` contains references to other acceleration structures, `dst` will reference the same acceleration structures.
- `VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR` serializes the acceleration structure to a semi-opaque format which can be reloaded on a compatible implementation.
- `VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR` deserializes the semi-opaque serialization format in the buffer to the acceleration structure.

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkCopyAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureInfoKHR.html), [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html), [VkCopyMemoryToAccelerationStructureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToAccelerationStructureInfoKHR.html), [vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyAccelerationStructureModeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0