# VK_EXT_opacity_micromap(3) Manual Page

## Name

VK_EXT_opacity_micromap - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

397

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)  
and  
     [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)  
     or  
     [Version 1.3](#versions-1.3)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_opacity_micromap](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_opacity_micromap.html)

## <a href="#_contact" class="anchor"></a>Contact

- Christoph Kubisch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_opacity_micromap%5D%20@pixeljetstream%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_opacity_micromap%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pixeljetstream</a>

- Eric Werness

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_opacity_micromap](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_opacity_micromap.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-08-24

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_opacity_micromap`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_opacity_micromap.txt)

**Contributors**  
- Christoph Kubisch, NVIDIA

- Eric Werness, NVIDIA

- Josh Barczak, Intel

- Stu Smith, AMD

## <a href="#_description" class="anchor"></a>Description

When adding transparency to a ray traced scene, an application can
choose between further tessellating the geometry or using an any-hit
shader to allow the ray through specific parts of the geometry. These
options have the downside of either significantly increasing memory
consumption or adding runtime overhead to run shader code in the middle
of traversal, respectively.

This extension adds the ability to add an *opacity micromap* to geometry
when building an acceleration structure. The opacity micromap compactly
encodes opacity information which can be read by the implementation to
mark parts of triangles as opaque or transparent. The format is
externally visible to allow the application to compress its internal
geometry and surface representations into the compressed format ahead of
time. The compressed format subdivides each triangle into a set of
subtriangles, each of which can be assigned either two or four opacity
values. These opacity values can control if a ray hitting that
subtriangle is treated as an opaque hit, complete miss, or possible hit,
depending on the controls described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-opacity-micromap"
target="_blank" rel="noopener">Ray Opacity Micromap</a>.

This extension provides:

- a [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) structure to store the micromap,

- functions similar to acceleration structure build functions to build
  the opacity micromap array, and

- a structure to extend
  [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)
  to attach a micromap to the geometry of the acceleration structure.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBuildMicromapsEXT.html)

- [vkCmdBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildMicromapsEXT.html)

- [vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToMicromapEXT.html)

- [vkCmdCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapEXT.html)

- [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html)

- [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteMicromapsPropertiesEXT.html)

- [vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToMicromapEXT.html)

- [vkCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapEXT.html)

- [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapToMemoryEXT.html)

- [vkCreateMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMicromapEXT.html)

- [vkDestroyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyMicromapEXT.html)

- [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMicromapCompatibilityEXT.html)

- [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html)

- [vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteMicromapsPropertiesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html)

- [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapInfoEXT.html)

- [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)

- [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)

- [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildSizesInfoEXT.html)

- [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)

- [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html)

- [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html)

- [VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapVersionInfoEXT.html)

- Extending
  [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html):

  - [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceOpacityMicromapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkBuildMicromapFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagBitsEXT.html)

- [VkBuildMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapModeEXT.html)

- [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html)

- [VkMicromapCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagBitsEXT.html)

- [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html)

- [VkOpacityMicromapFormatEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpacityMicromapFormatEXT.html)

- [VkOpacityMicromapSpecialIndexEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpacityMicromapSpecialIndexEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkBuildMicromapFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagsEXT.html)

- [VkMicromapCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_OPACITY_MICROMAP_EXTENSION_NAME`

- `VK_EXT_OPACITY_MICROMAP_SPEC_VERSION`

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_MICROMAP_READ_BIT_EXT`

  - `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`

- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html):

  - `VK_BUFFER_USAGE_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT`

  - `VK_BUFFER_USAGE_MICROMAP_STORAGE_BIT_EXT`

- Extending
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISABLE_OPACITY_MICROMAPS_EXT`

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_DATA_UPDATE_EXT`

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_OPACITY_MICROMAP_UPDATE_EXT`

- Extending
  [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagBitsKHR.html):

  - `VK_GEOMETRY_INSTANCE_DISABLE_OPACITY_MICROMAPS_EXT`

  - `VK_GEOMETRY_INSTANCE_FORCE_OPACITY_MICROMAP_2_STATE_EXT`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_MICROMAP_EXT`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT`

  - `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_TRIANGLES_OPACITY_MICROMAP_EXT`

  - `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_MICROMAP_INFO_EXT`

  - `VK_STRUCTURE_TYPE_COPY_MICROMAP_INFO_EXT`

  - `VK_STRUCTURE_TYPE_COPY_MICROMAP_TO_MEMORY_INFO_EXT`

  - `VK_STRUCTURE_TYPE_MICROMAP_BUILD_INFO_EXT`

  - `VK_STRUCTURE_TYPE_MICROMAP_BUILD_SIZES_INFO_EXT`

  - `VK_STRUCTURE_TYPE_MICROMAP_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_MICROMAP_VERSION_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPACITY_MICROMAP_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPACITY_MICROMAP_PROPERTIES_EXT`

## <a href="#_reference_code" class="anchor"></a>Reference Code

``` c
uint32_t BarycentricsToSpaceFillingCurveIndex(float u, float v, uint32_t level)
{
    u = clamp(u, 0.0f, 1.0f);
    v = clamp(v, 0.0f, 1.0f);

    uint32_t iu, iv, iw;

    // Quantize barycentric coordinates
    float fu = u * (1u << level);
    float fv = v * (1u << level);

    iu = (uint32_t)fu;
    iv = (uint32_t)fv;

    float uf = fu - float(iu);
    float vf = fv - float(iv);

    if (iu >= (1u << level)) iu = (1u << level) - 1u;
    if (iv >= (1u << level)) iv = (1u << level) - 1u;

    uint32_t iuv = iu + iv;

    if (iuv >= (1u << level))
        iu -= iuv - (1u << level) + 1u;

    iw = ~(iu + iv);

    if (uf + vf >= 1.0f && iuv < (1u << level) - 1u) --iw;

    uint32_t b0 = ~(iu ^ iw);
    b0 &= ((1u << level) - 1u);
    uint32_t t = (iu ^ iv) & b0;

    uint32_t f = t;
    f ^= f >> 1u;
    f ^= f >> 2u;
    f ^= f >> 4u;
    f ^= f >> 8u;
    uint32_t b1 = ((f ^ iu) & ~b0) | t;

    // Interleave bits
    b0 = (b0 | (b0 << 8u)) & 0x00ff00ffu;
    b0 = (b0 | (b0 << 4u)) & 0x0f0f0f0fu;
    b0 = (b0 | (b0 << 2u)) & 0x33333333u;
    b0 = (b0 | (b0 << 1u)) & 0x55555555u;
    b1 = (b1 | (b1 << 8u)) & 0x00ff00ffu;
    b1 = (b1 | (b1 << 4u)) & 0x0f0f0f0fu;
    b1 = (b1 | (b1 << 2u)) & 0x33333333u;
    b1 = (b1 | (b1 << 1u)) & 0x55555555u;

    return b0 | (b1 << 1u);
}
```

## <a href="#_issues" class="anchor"></a>Issues

\(1\) Is the build actually similar to an acceleration structure build?

- Resolved: The build should be much lighter-weight than an acceleration
  structure build, but the infrastructure is similar enough that it
  makes sense to keep the concepts compatible.

\(2\) Why does VkMicromapUsageEXT not have type/pNext?

- Resolved: There can be a very large number of these structures, so
  doubling the size of these can be significant memory consumption.
  Also, an application may be loading these directly from a file which
  is more compatible with it being a flat structure. The including
  structures are extensible and are probably a more suitable place to
  add extensibility.

\(3\) Why is there a SPIR-V extension?

- Resolved: There is a ray flag. To be consistent with how the existing
  ray tracing extensions work that ray flag needs its own extension.

\(4\) Should there be indirect micromap build?

- Resolved: Not for now. There is more in-depth usage metadata required
  and it seems less likely that something like a GPU culling system
  would need to change the counts for a micromap.

\(5\) Should micromaps have a micromap device address?

- Resolved: There is no need right now (can just use the handle) but
  that is a bit different from acceleration structures, though the two
  are not completely parallel in their usage.

\(6\) Why are the alignment requirements defined as a mix of hardcoded
values and caps?

- Resolved: This is most parallel with the definition of
  [`VK_KHR_acceleration_structure`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html)
  and maintaining commonality makes it easier for applications to share
  memory.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2022-06-22 (Eric Werness)

  - EXTify and clean up for discussion

- Revision 1, 2022-01-01 (Eric Werness)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_opacity_micromap"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
