# VK_NV_ray_tracing_validation(3) Manual Page

## Name

VK_NV_ray_tracing_validation - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

569

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Vikram Kushwaha <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_ray_tracing_validation%5D%20@vkushwaha-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_ray_tracing_validation%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>vkushwaha-nv</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_NV_ray_tracing_validation](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_ray_tracing_validation.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-03-04

**Contributors**  
- Vikram Kushwaha, NVIDIA

- Eric Werness, NVIDIA

- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for performing ray tracing validation at an
implementation level.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRayTracingValidationFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingValidationFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_RAY_TRACING_VALIDATION_EXTENSION_NAME`

- `VK_NV_RAY_TRACING_VALIDATION_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_VALIDATION_FEATURES_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-03-04 (Vikram Kushwaha)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_ray_tracing_validation"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
