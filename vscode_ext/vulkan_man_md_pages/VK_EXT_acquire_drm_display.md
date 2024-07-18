# VK_EXT_acquire_drm_display(3) Manual Page

## Name

VK_EXT_acquire_drm_display - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

286

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_direct_mode_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_direct_mode_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Drew DeVault <sir@cmpwn.com>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-06-09

**IP Status**  
No known IP claims.

**Contributors**  
- Simon Zeni, Status Holdings, Ltd.

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to take exclusive control of a
display using the Direct Rendering Manager (DRM) interface. When
acquired, the display will be under full control of the application
until the display is either released or the connector is unplugged.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkAcquireDrmDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireDrmDisplayEXT.html)

- [vkGetDrmDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDrmDisplayEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_ACQUIRE_DRM_DISPLAY_EXTENSION_NAME`

- `VK_EXT_ACQUIRE_DRM_DISPLAY_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-05-11 (Simon Zeni)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_acquire_drm_display"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
