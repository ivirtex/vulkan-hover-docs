# VK_NV_displacement_micromap(3) Manual Page

## Name

VK_NV_displacement_micromap - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

398

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html)  

- **This is a *provisional* extension and **must** be used with caution.
  See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#boilerplate-provisional-header"
  target="_blank" rel="noopener">description</a> of provisional header
  files for enablement and stability details.**

## <a href="#_contact" class="anchor"></a>Contact

- Christoph Kubisch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_displacement_micromap%5D%20@pixeljetstream%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_displacement_micromap%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pixeljetstream</a>

- Eric Werness <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_displacement_micromap%5D%20@ewerness-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_displacement_micromap%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ewerness-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-17

**Interactions and External Dependencies**  
TBD

**Contributors**  
- Christoph Kubisch, NVIDIA

- Eric Werness, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Ray tracing can very efficiently render from geometry which has very
fine detail, but when using only a basic triangle representation, memory
consumption can be an issue. This extension adds the ability to add a
*displacement map* to add more detail to triangles in an acceleration
structure with an efficient in-memory format. The format is externally
visible to allow the application to compress its internal geometry
representations into the compressed format ahead of time. This format
adds displacements along a defined vector to subtriangle vertices which
are subdivided from the main triangles.

This extension provides:

- a new [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) format for the
  displacement micromap,

- a structure to extend
  [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)
  to attach a displacement micromap to the geometry of the acceleration
  structure,

- enums extending
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html)
  to allow for updates.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html):

  - [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDisplacementMicromapFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDisplacementMicromapFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDisplacementMicromapFormatNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplacementMicromapFormatNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_DISPLACEMENT_MICROMAP_EXTENSION_NAME`

- `VK_NV_DISPLACEMENT_MICROMAP_SPEC_VERSION`

- Extending
  [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):

  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISPLACEMENT_MICROMAP_UPDATE_NV`

- Extending [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html):

  - `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_TRIANGLES_DISPLACEMENT_MICROMAP_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISPLACEMENT_MICROMAP_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISPLACEMENT_MICROMAP_PROPERTIES_NV`

## <a href="#_issues" class="anchor"></a>Issues

\(1\) What is the status of this extension?

- Provisional and expected to change. The broad structure and encoding
  format are stable, but there will likely be changes to the structures,
  enumerant values, and shader interface.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-03-17 (Eric Werness)

  - Initial public revision

- Revision 2, 2023-07-07 (Eric Werness)

  - Add shader support for decode intrinsics

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_displacement_micromap"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
