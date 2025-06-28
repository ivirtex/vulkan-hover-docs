# VK\_KHR\_uniform\_buffer\_standard\_layout(3) Manual Page

## Name

VK\_KHR\_uniform\_buffer\_standard\_layout - device extension



## [](#_registered_extension_number)Registered Extension Number

254

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Graeme Leese [\[GitHub\]gnl21](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_uniform_buffer_standard_layout%5D%20%40gnl21%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_uniform_buffer_standard_layout%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-25

**Contributors**

- Graeme Leese, Broadcom
- Jeff Bolz, NVIDIA
- Tobias Hector, AMD
- Faith Ekstrand, Intel
- Neil Henning, AMD

## [](#_description)Description

This extension enables tighter array and structure packing to be used with uniform buffers.

It modifies the alignment rules for uniform buffers, allowing for tighter packing of arrays and structures. This allows, for example, the std430 layout, as defined in [GLSL](https://registry.khronos.org/OpenGL/specs/gl/GLSLangSpec.4.60.pdf) to be supported in uniform buffers.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME`
- `VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2019-01-25 (Graeme Leese)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_uniform_buffer_standard_layout)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0