# VK\_NV\_ray\_tracing\_validation(3) Manual Page

## Name

VK\_NV\_ray\_tracing\_validation - device extension



## [](#_registered_extension_number)Registered Extension Number

569

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing_validation%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing_validation%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_ray\_tracing\_validation](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_ray_tracing_validation.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-03-04

**Contributors**

- Vikram Kushwaha, NVIDIA
- Eric Werness, NVIDIA
- Piers Daniell, NVIDIA

## [](#_description)Description

This extension adds support for performing ray tracing validation at an implementation level.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRayTracingValidationFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingValidationFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_RAY_TRACING_VALIDATION_EXTENSION_NAME`
- `VK_NV_RAY_TRACING_VALIDATION_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_VALIDATION_FEATURES_NV`

## [](#_version_history)Version History

- Revision 1, 2024-03-04 (Vikram Kushwaha)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_ray_tracing_validation).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0