# VK_EXT_validation_features(3) Manual Page

## Name

VK_EXT_validation_features - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

248

## <a href="#_revision" class="anchor"></a>Revision

6

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by [VK_EXT_layer_settings](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_layer_settings.html)
  extension

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Debugging tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Karl Schultz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_validation_features%5D%20@karl-lunarg%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_validation_features%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>karl-lunarg</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension provides the
[VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html) struct that can
be included in the `pNext` chain of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure passed as
the `pCreateInfo` parameter of
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html). The structure contains an
array of
[VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureEnableEXT.html) enum
values that enable specific validation features that are disabled by
default. The structure also contains an array of
[VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureDisableEXT.html) enum
values that disable specific validation layer features that are enabled
by default.

## <a href="#_deprecation_by_vk_ext_layer_settings" class="anchor"></a>Deprecation by `VK_EXT_layer_settings`

Functionality in this extension is subsumed into the
[`VK_EXT_layer_settings`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_layer_settings.html) extension.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html):

  - [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureDisableEXT.html)

- [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureEnableEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_VALIDATION_FEATURES_EXTENSION_NAME`

- `VK_EXT_VALIDATION_FEATURES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VALIDATION_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

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

  - Marked as deprecated by VK_EXT_layer_settings

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_validation_features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
