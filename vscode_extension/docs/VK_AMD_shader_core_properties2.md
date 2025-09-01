# VK\_AMD\_shader\_core\_properties2(3) Manual Page

## Name

VK\_AMD\_shader\_core\_properties2 - device extension



## [](#_registered_extension_number)Registered Extension Number

228

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_AMD\_shader\_core\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_core_properties.html)

## [](#_contact)Contact

- Matthaeus G. Chajdas [\[GitHub\]anteru](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_core_properties2%5D%20%40anteru%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_core_properties2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-07-26

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Tobias Hector, AMD

## [](#_description)Description

This extension exposes additional shader core properties for a target physical device through the `VK_KHR_get_physical_device_properties2` extension.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderCoreProperties2AMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderCoreProperties2AMD.html)

## [](#_new_enums)New Enums

- [VkShaderCorePropertiesFlagBitsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCorePropertiesFlagBitsAMD.html)

## [](#_new_bitmasks)New Bitmasks

- [VkShaderCorePropertiesFlagsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCorePropertiesFlagsAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_SHADER_CORE_PROPERTIES_2_EXTENSION_NAME`
- `VK_AMD_SHADER_CORE_PROPERTIES_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD`

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2019-07-26 (Matthaeus G. Chajdas)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_shader_core_properties2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0