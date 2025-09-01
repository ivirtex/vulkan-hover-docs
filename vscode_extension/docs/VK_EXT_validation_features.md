# VK\_EXT\_validation\_features(3) Manual Page

## Name

VK\_EXT\_validation\_features - instance extension



## [](#_registered_extension_number)Registered Extension Number

248

## [](#_revision)Revision

6

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_EXT\_layer\_settings](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_layer_settings.html) extension

## [](#_special_use)Special Use

- [Debugging tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Karl Schultz [\[GitHub\]karl-lunarg](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_validation_features%5D%20%40karl-lunarg%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_validation_features%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-11-14

**IP Status**

No known IP claims.

**Contributors**

- Karl Schultz, LunarG
- Dave Houlton, LunarG
- Mark Lobodzinski, LunarG
- Camden Stocker, LunarG
- Tony Barbour, LunarG
- John Zulauf, LunarG

## [](#_description)Description

This extension provides the [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html) structure that can be included in the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure passed as the `pCreateInfo` parameter of [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html). The structure contains an array of [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureEnableEXT.html) enum values that enable specific validation features that are disabled by default. The structure also contains an array of [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureDisableEXT.html) enum values that disable specific validation layer features that are enabled by default.

## [](#_deprecation_by_vk_ext_layer_settings)Deprecation by `VK_EXT_layer_settings`

Functionality in this extension is subsumed into the `VK_EXT_layer_settings` extension.

## [](#_new_structures)New Structures

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html), [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html):
  
  - [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)

## [](#_new_enums)New Enums

- [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureDisableEXT.html)
- [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureEnableEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_VALIDATION_FEATURES_EXTENSION_NAME`
- `VK_EXT_VALIDATION_FEATURES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VALIDATION_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2018-11-14 (Karl Schultz)
  
  - Initial revision
- Revision 2, 2019-08-06 (Mark Lobodzinski)
  
  - Add Best Practices enable
- Revision 3, 2020-03-04 (Tony Barbour)
  
  - Add Debug Printf enable
- Revision 4, 2020-07-29 (John Zulauf)
  
  - Add Synchronization Validation enable
- Revision 5, 2021-05-18 (Tony Barbour)
  
  - Add Shader Validation Cache disable
- Revision 6, 2023-09-25 (Christophe Riccio)
  
  - Marked as deprecated by VK\_EXT\_layer\_settings

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_validation_features).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0