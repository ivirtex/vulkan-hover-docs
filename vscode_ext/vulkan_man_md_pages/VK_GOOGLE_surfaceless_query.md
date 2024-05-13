# VK_GOOGLE_surfaceless_query(3) Manual Page

## Name

VK_GOOGLE_surfaceless_query - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

434

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GOOGLE_surfaceless_query%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GOOGLE_surfaceless_query%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_GOOGLE_surfaceless_query](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_GOOGLE_surfaceless_query.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-08-03

**IP Status**  
No known IP claims.

**Contributors**  
- Ian Elliott, Google

- Shahbaz Youssefi, Google

- James Jones, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows the
[vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html)
and
[vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)
functions to accept [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) as their
`surface` parameter, allowing potential surface formats, color spaces
and present modes to be queried without providing a surface.
Identically,
[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html),
[vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html),
and
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)
would accept [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) in
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface`.
**This can only be supported on platforms where the results of these
queries are surface-agnostic and a single presentation engine is the
implicit target of all present operations**.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_GOOGLE_SURFACELESS_QUERY_EXTENSION_NAME`

- `VK_GOOGLE_SURFACELESS_QUERY_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-12-14 (Shahbaz Youssefi)

  - Internal revisions

- Revision 2, 2022-08-03 (Shahbaz Youssefi)

  - Precisions to which parts of the query responses are defined when
    surfaceless

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_GOOGLE_surfaceless_query"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
