# VK_FUCHSIA_imagepipe_surface(3) Manual Page

## Name

VK_FUCHSIA_imagepipe_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

215

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Craig Stout <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_imagepipe_surface%5D%20@cdotstout%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_imagepipe_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cdotstout</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-07-27

**IP Status**  
No known IP claims.

**Contributors**  
- Craig Stout, Google

- Ian Elliott, Google

- Jesse Hall, Google

## <a href="#_description" class="anchor"></a>Description

The `VK_FUCHSIA_imagepipe_surface` extension is an instance extension.
It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)
object (defined by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)
extension) that refers to a Fuchsia `imagePipeHandle`.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImagePipeSurfaceFUCHSIA.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkImagePipeSurfaceCreateFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateFlagsFUCHSIA.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_FUCHSIA_IMAGEPIPE_SURFACE_EXTENSION_NAME`

- `VK_FUCHSIA_IMAGEPIPE_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-07-27 (Craig Stout)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_FUCHSIA_imagepipe_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
