# VK_AMD_negative_viewport_height(3) Manual Page

## Name

VK_AMD_negative_viewport_height - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

36

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Obsoleted* by [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html)
  extension

  - Which in turn was *promoted* to <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
    target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Matthaeus G. Chajdas <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_negative_viewport_height%5D%20@anteru%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_negative_viewport_height%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>anteru</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-09-02

**IP Status**  
No known IP claims.

**Contributors**  
- Matthaeus G. Chajdas, AMD

- Graham Sellers, AMD

- Baldur Karlsson

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to specify a negative viewport
height. The result is that the viewport transformation will flip along
the y-axis.

- Allow negative height to be specified in the
  [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)::`height` field to perform y-inversion
  of the clip-space to framebuffer-space transform. This allows apps to
  avoid having to use `gl_Position.y = -gl_Position.y` in shaders also
  targeting other APIs.

## <a href="#_obsoletion_by_vk_khr_maintenance1_and_vulkan_1_1"
class="anchor"></a>Obsoletion by `VK_KHR_maintenance1` and Vulkan 1.1

Functionality in this extension is included in `VK_KHR_maintenance1` and
subsequently Vulkan 1.1. Due to some slight behavioral differences, this
extension **must** not be enabled alongside `VK_KHR_maintenance1`, or in
an instance created with version 1.1 or later requested in
[VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion`.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_NEGATIVE_VIEWPORT_HEIGHT_EXTENSION_NAME`

- `VK_AMD_NEGATIVE_VIEWPORT_HEIGHT_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-09-02 (Matthaeus Chajdas)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_negative_viewport_height"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
