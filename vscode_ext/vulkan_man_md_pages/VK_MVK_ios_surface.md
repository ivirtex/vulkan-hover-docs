# VK_MVK_ios_surface(3) Manual Page

## Name

VK_MVK_ios_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

123

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by [VK_EXT_metal_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_surface.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Bill Hollings <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_MVK_ios_surface%5D%20@billhollings%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_MVK_ios_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>billhollings</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-07-31

**IP Status**  
No known IP claims.

**Contributors**  
- Bill Hollings, The Brenwill Workshop Ltd.

## <a href="#_description" class="anchor"></a>Description

The `VK_MVK_ios_surface` extension is an instance extension. It provides
a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) object
(defined by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) extension) based
on a `UIView`, the native surface type of iOS, which is underpinned by a
[CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html), to support rendering to the surface
using Apple’s Metal framework.

## <a href="#_deprecation_by_vk_ext_metal_surface" class="anchor"></a>Deprecation by `VK_EXT_metal_surface`

The `VK_MVK_ios_surface` extension is considered deprecated and has been
superseded by the [`VK_EXT_metal_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_surface.html)
extension.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIOSSurfaceMVK.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateInfoMVK.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkIOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateFlagsMVK.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_MVK_IOS_SURFACE_EXTENSION_NAME`

- `VK_MVK_IOS_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-15 (Bill Hollings)

  - Initial draft.

- Revision 2, 2017-02-24 (Bill Hollings)

  - Minor syntax fix to emphasize firm requirement for `UIView` to be
    backed by a `CAMetalLayer`.

- Revision 3, 2020-07-31 (Bill Hollings)

  - Update documentation on requirements for `UIView`.

  - Mark as deprecated by `VK_EXT_metal_surface`.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MVK_ios_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
