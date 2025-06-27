# VK_EXT_validation_flags(3) Manual Page

## Name

VK_EXT_validation_flags - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

62

## <a href="#_revision" class="anchor"></a>Revision

3

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

- Tobin Ehlis <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_validation_flags%5D%20@tobine%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_validation_flags%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobine</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-08-19

**IP Status**  
No known IP claims.

**Contributors**  
- Tobin Ehlis, Google

- Courtney Goeltzenleuchter, Google

## <a href="#_description" class="anchor"></a>Description

This extension provides the
[VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFlagsEXT.html) struct that can be
included in the `pNext` chain of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure passed as
the `pCreateInfo` parameter of
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html). The structure contains an
array of [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCheckEXT.html) values that
will be disabled by the validation layers.

## <a href="#_deprecation_by_vk_ext_layer_settings" class="anchor"></a>Deprecation by `VK_EXT_layer_settings`

Functionality in this extension is subsumed into the
[`VK_EXT_layer_settings`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_layer_settings.html) extension.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html):

  - [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFlagsEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCheckEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_VALIDATION_FLAGS_EXTENSION_NAME`

- `VK_EXT_VALIDATION_FLAGS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 3, 2023-09-25 (Christophe Riccio)

  - Marked as deprecated by VK_EXT_layer_settings

- Revision 2, 2019-08-19 (Mark Lobodzinski)

  - Marked as deprecated by VK_EXT_validation_features

- Revision 1, 2016-08-26 (Courtney Goeltzenleuchter)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_validation_flags"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
