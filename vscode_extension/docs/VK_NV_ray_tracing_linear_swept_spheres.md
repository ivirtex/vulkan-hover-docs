# VK\_NV\_ray\_tracing\_linear\_swept\_spheres(3) Manual Page

## Name

VK\_NV\_ray\_tracing\_linear\_swept\_spheres - device extension



## [](#_registered_extension_number)Registered Extension Number

430

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_linear\_swept\_spheres](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_linear_swept_spheres.html)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing_linear_swept_spheres%5D%20%40vkushwaha%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing_linear_swept_spheres%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_ray\_tracing\_linear\_swept\_spheres](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_ray_tracing_linear_swept_spheres.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-03

**Interactions and External Dependencies**

- This extension requires [`SPV_NV_linear_swept_spheres`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_linear_swept_spheres.html)
- This extension provides API support for [`GL_NV_linear_swept_spheres`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_linear_swept_spheres.txt)

**Contributors**

- Vikram Kushwaha, NVIDIA
- Eric Werness, NVIDIA
- Daniel Koch, NVIDIA
- Ashwin Lele, NVIDIA
- Nathan Morrical, NVIDIA

## [](#_description)Description

This extension adds two new primitives for ray tracing: a sphere primitive and a linear swept sphere (LSS) primitive. The purpose of the LSS primitive is to enable rendering of high quality hair and fur using a compact primitive representation encoded in the acceleration structure. Sphere primitives are defined by a position and a radius and are a subset of LSS, but are useful in their own right, for example for particle systems.

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_NV_linear_swept_spheres`

## [](#_new_structures)New Structures

- Extending [VkAccelerationStructureGeometryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryKHR.html):
  
  - [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)
  - [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV.html)

## [](#_new_enums)New Enums

- [VkRayTracingLssIndexingModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssIndexingModeNV.html)
- [VkRayTracingLssPrimitiveEndCapsModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingLssPrimitiveEndCapsModeNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_RAY_TRACING_LINEAR_SWEPT_SPHERES_EXTENSION_NAME`
- `VK_NV_RAY_TRACING_LINEAR_SWEPT_SPHERES_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_ACCELERATION_STRUCTURE_RADIUS_BUFFER_BIT_NV`
- Extending [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html):
  
  - `VK_GEOMETRY_TYPE_LINEAR_SWEPT_SPHERES_NV`
  - `VK_GEOMETRY_TYPE_SPHERES_NV`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_RAY_TRACING_ALLOW_SPHERES_AND_LINEAR_SWEPT_SPHERES_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_LINEAR_SWEPT_SPHERES_DATA_NV`
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_SPHERES_DATA_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_LINEAR_SWEPT_SPHERES_FEATURES_NV`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`HitIsSphereNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitissphere)
- [`HitIsLSSNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitislss)
- [`HitSpherePositionNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitsphereposition)
- [`HitSphereRadiusNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitsphereradius)
- [`HitLSSPositionsNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitlsspositions)
- [`HitLSSRadiiNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-hitlssradii)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`RayTracingSpheresGeometryNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTracingSpheresGeometryNV)
- [`RayTracingLinearSweptSpheresGeometryNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTracingLinearSweptSpheresGeometryNV)

## [](#_version_history)Version History

- Revision 1, 2025-01-03 (Vikram Kushwaha)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_ray_tracing_linear_swept_spheres).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0