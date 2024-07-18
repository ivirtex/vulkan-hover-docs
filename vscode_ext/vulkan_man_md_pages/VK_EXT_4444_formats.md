# VK_EXT_4444_formats(3) Manual Page

## Name

VK_EXT_4444_formats - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

341

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Joshua Ashton <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_4444_formats%5D%20@Joshua-Ashton%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_4444_formats%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>Joshua-Ashton</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-07-28

**IP Status**  
No known IP claims.

**Contributors**  
- Joshua Ashton, Valve

- Faith Ekstrand, Intel

## <a href="#_description" class="anchor"></a>Description

This extension defines the `VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT` and
`VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT` formats which are defined in other
current graphics APIs.

This extension may be useful for building translation layers for those
APIs or for porting applications that use these formats without having
to resort to swizzles.

When VK_EXT_custom_border_color is used, these formats are not subject
to the same restrictions for border color without format as with
VK_FORMAT_B4G4R4A4_UNORM_PACK16.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevice4444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice4444FormatsFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_4444_FORMATS_EXTENSION_NAME`

- `VK_EXT_4444_FORMATS_SPEC_VERSION`

- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html):

  - `VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT`

  - `VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_4444_FORMATS_FEATURES_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

This extension has been partially promoted. The format enumerants
introduced by the extension are included in core Vulkan 1.3, with the
EXT suffix omitted. However, runtime support for these formats is
optional in core Vulkan 1.3, while if this extension is supported,
runtime support is mandatory. The feature structure is not promoted. The
original enum names are still available as aliases of the core
functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-07-04 (Joshua Ashton)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_4444_formats"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
