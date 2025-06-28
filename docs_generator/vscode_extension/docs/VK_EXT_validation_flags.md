# VK\_EXT\_validation\_flags(3) Manual Page

## Name

VK\_EXT\_validation\_flags - instance extension



## [](#_registered_extension_number)Registered Extension Number

62

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_EXT\_layer\_settings](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_layer_settings.html) extension

## [](#_special_use)Special Use

- [Debugging tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Tobin Ehlis [\[GitHub\]tobine](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_validation_flags%5D%20%40tobine%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_validation_flags%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-08-19

**IP Status**

No known IP claims.

**Contributors**

- Tobin Ehlis, Google
- Courtney Goeltzenleuchter, Google

## [](#_description)Description

This extension provides the [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html) structure that can be included in the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure passed as the `pCreateInfo` parameter of [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html). The structure contains an array of [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCheckEXT.html) values that will be disabled by the validation layers.

## [](#_deprecation_by_vk_ext_layer_settings)Deprecation by `VK_EXT_layer_settings`

Functionality in this extension is subsumed into the `VK_EXT_layer_settings` extension.

## [](#_new_structures)New Structures

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html):
  
  - [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html)

## [](#_new_enums)New Enums

- [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCheckEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_VALIDATION_FLAGS_EXTENSION_NAME`
- `VK_EXT_VALIDATION_FLAGS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT`

## [](#_version_history)Version History

- Revision 3, 2023-09-25 (Christophe Riccio)
  
  - Marked as deprecated by VK\_EXT\_layer\_settings
- Revision 2, 2019-08-19 (Mark Lobodzinski)
  
  - Marked as deprecated by VK\_EXT\_validation\_features
- Revision 1, 2016-08-26 (Courtney Goeltzenleuchter)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_validation_flags)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0