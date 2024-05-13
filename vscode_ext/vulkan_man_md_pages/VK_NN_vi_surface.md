# VK_NN_vi_surface(3) Manual Page

## Name

VK_NN_vi_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

63

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Mathias Heyer mheyer

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-12-02

**IP Status**  
No known IP claims.

**Contributors**  
- Mathias Heyer, NVIDIA

- Michael Chock, NVIDIA

- Yasuhiro Yoshioka, Nintendo

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

The `VK_NN_vi_surface` extension is an instance extension. It provides a
mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) object (defined
by the [`VK_KHR_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) extension) associated
with an `nn`::`vi`::`Layer`.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateViSurfaceNN.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateInfoNN.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkViSurfaceCreateFlagsNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateFlagsNN.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NN_VI_SURFACE_EXTENSION_NAME`

- `VK_NN_VI_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does VI need a way to query for compatibility between a particular
physical device (and queue family?) and a specific VI display?

**RESOLVED**: No. It is currently always assumed that the device and
display will always be compatible.

2\) [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateInfoNN.html)::`pWindow`
is intended to store an `nn`::`vi`::`NativeWindowHandle`, but its
declared type is a bare `void*` to store the window handle. Why the
discrepancy?

**RESOLVED**: It is for C compatibility. The definition for the VI
native window handle type is defined inside the `nn`::`vi` C++
namespace. This prevents its use in C source files.
`nn`::`vi`::`NativeWindowHandle` is always defined to be `void*`, so
this extension uses `void*` to match.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-2 (Michael Chock)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NN_vi_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
