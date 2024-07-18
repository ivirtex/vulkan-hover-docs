# VkBuildAccelerationStructureFlagBitsKHR(3) Manual Page

## Name

VkBuildAccelerationStructureFlagBitsKHR - Bitmask specifying additional
parameters for acceleration structure builds



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)::`flags`
or
[VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)::`flags`
specifying additional parameters for acceleration structure builds, are:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkBuildAccelerationStructureFlagBitsKHR {
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR = 0x00000001,
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR = 0x00000002,
    VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR = 0x00000004,
    VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR = 0x00000008,
    VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR = 0x00000010,
  // Provided by VK_NV_ray_tracing_motion_blur
    VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV = 0x00000020,
  // Provided by VK_EXT_opacity_micromap
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_UPDATE_EXT = 0x00000040,
  // Provided by VK_EXT_opacity_micromap
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISABLE_OPACITY_MICROMAPS_EXT = 0x00000080,
  // Provided by VK_EXT_opacity_micromap
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_DATA_UPDATE_EXT = 0x00000100,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_NV_displacement_micromap
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISPLACEMENT_MICROMAP_UPDATE_NV = 0x00000200,
#endif
  // Provided by VK_KHR_ray_tracing_position_fetch
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR = 0x00000800,
  // Provided by VK_NV_ray_tracing
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV = VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_NV = VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_NV = VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_NV = VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_NV = VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR,
} VkBuildAccelerationStructureFlagBitsKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkBuildAccelerationStructureFlagBitsKHR VkBuildAccelerationStructureFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR` indicates that
  the specified acceleration structure **can** be updated with a `mode`
  of `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` in
  [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)
  or an `update` of `VK_TRUE` in
  [vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructureNV.html)
  .

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` indicates
  that the specified acceleration structure **can** act as the source
  for a copy acceleration structure command with `mode` of
  `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR` to produce a
  compacted acceleration structure.

- `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_TRACE_BIT_KHR` indicates
  that the given acceleration structure build **should** prioritize
  trace performance over build time.

- `VK_BUILD_ACCELERATION_STRUCTURE_PREFER_FAST_BUILD_BIT_KHR` indicates
  that the given acceleration structure build **should** prioritize
  build time over trace performance.

- `VK_BUILD_ACCELERATION_STRUCTURE_LOW_MEMORY_BIT_KHR` indicates that
  this acceleration structure **should** minimize the size of the
  scratch memory and the final result acceleration structure,
  potentially at the expense of build time or trace performance.

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_UPDATE_EXT`
  indicates that the opacity micromaps associated with the specified
  acceleration structure **may** change with an acceleration structure
  update.

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_DATA_UPDATE_EXT`
  indicates that the data of the opacity micromaps associated with the
  specified acceleration structure **may** change with an acceleration
  structure update.

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISABLE_OPACITY_MICROMAPS_EXT`
  indicates that the specified acceleration structure **may** be
  referenced in an instance with
  `VK_GEOMETRY_INSTANCE_DISABLE_OPACITY_MICROMAPS_EXT` set.

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR` indicates that
  the specified acceleration structure **can** be used when fetching the
  vertex positions of a hit triangle.

- `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISPLACEMENT_MICROMAP_UPDATE_NV`
  indicates that the displacement micromaps associated with the
  specified acceleration structure **may** change with an acceleration
  structure update.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_KHR</code> and
<code>VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR</code>
<strong>may</strong> take more time and memory than a normal build, and
so <strong>should</strong> only be used when those features are
needed.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBuildAccelerationStructureFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
