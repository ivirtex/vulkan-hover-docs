# VK\_NV\_displacement\_micromap(3) Manual Page

## Name

VK\_NV\_displacement\_micromap - device extension



## [](#_registered_extension_number)Registered Extension Number

398

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html)

- **This is a *provisional* extension and must be used with caution. See the [description](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-provisional-header) of provisional header files for enablement and stability details.**

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html) extension

## [](#_contact)Contact

- Christoph Kubisch [\[GitHub\]pixeljetstream](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_displacement_micromap%5D%20%40pixeljetstream%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_displacement_micromap%20extension%2A)
- Eric Werness [\[GitHub\]ewerness-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_displacement_micromap%5D%20%40ewerness-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_displacement_micromap%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-17

**Interactions and External Dependencies**

TBD

**Contributors**

- Christoph Kubisch, NVIDIA
- Eric Werness, NVIDIA

## [](#_description)Description

Ray tracing can very efficiently render from geometry which has very fine detail, but when using only a basic triangle representation, memory consumption can be an issue. This extension adds the ability to add a *displacement map* to add more detail to triangles in an acceleration structure with an efficient in-memory format. The format is externally visible to allow the application to compress its internal geometry representations into the compressed format ahead of time. This format adds displacements along a defined vector to subtriangle vertices which are subdivided from the main triangles.

This extension provides:

- a new [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html) format for the displacement micromap,
- a structure to extend [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html) to attach a displacement micromap to the geometry of the acceleration structure,
- enums extending [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html) to allow for updates.

## [](#_new_structures)New Structures

- Extending [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html):
  
  - [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDisplacementMicromapFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDisplacementMicromapFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkDisplacementMicromapFormatNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplacementMicromapFormatNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DISPLACEMENT_MICROMAP_EXTENSION_NAME`
- `VK_NV_DISPLACEMENT_MICROMAP_SPEC_VERSION`
- Extending [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):
  
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISPLACEMENT_MICROMAP_UPDATE_BIT_NV`
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DISPLACEMENT_MICROMAP_UPDATE_NV`
- Extending [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTypeEXT.html):
  
  - `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_TRIANGLES_DISPLACEMENT_MICROMAP_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISPLACEMENT_MICROMAP_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISPLACEMENT_MICROMAP_PROPERTIES_NV`

## [](#_issues)Issues

(1) What is the status of this extension?

- Deprecated. The VK\_NV\_cluster\_acceleration\_structure extension is not a one-to-one replacement for this extension but enables similar performance improvements for high-tessellation geometry and is considered the preferred direction to improve high-tessellation geometry performance.

## [](#_version_history)Version History

- Revision 1, 2023-03-17 (Eric Werness)
  
  - Initial public revision
- Revision 2, 2023-07-07 (Eric Werness)
  
  - Add shader support for decode intrinsics

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_displacement_micromap).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0