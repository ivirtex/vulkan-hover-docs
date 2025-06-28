# VK\_GOOGLE\_surfaceless\_query(3) Manual Page

## Name

VK\_GOOGLE\_surfaceless\_query - instance extension



## [](#_registered_extension_number)Registered Extension Number

434

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GOOGLE_surfaceless_query%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GOOGLE_surfaceless_query%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_GOOGLE\_surfaceless\_query](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_GOOGLE_surfaceless_query.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-08-03

**IP Status**

No known IP claims.

**Contributors**

- Ian Elliott, Google
- Shahbaz Youssefi, Google
- James Jones, NVIDIA

## [](#_description)Description

This extension allows the [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html) and [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html) functions to accept [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) as their `surface` parameter, allowing potential surface formats, color spaces and present modes to be queried without providing a surface. Identically, [vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html), [vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html), and [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html) would accept [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) in [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface`. **This can only be supported on platforms where the results of these queries are surface-agnostic and a single presentation engine is the implicit target of all present operations**.

## [](#_new_enum_constants)New Enum Constants

- `VK_GOOGLE_SURFACELESS_QUERY_EXTENSION_NAME`
- `VK_GOOGLE_SURFACELESS_QUERY_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 1, 2021-12-14 (Shahbaz Youssefi)
  
  - Internal revisions
- Revision 2, 2022-08-03 (Shahbaz Youssefi)
  
  - Precisions to which parts of the query responses are defined when surfaceless

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_GOOGLE_surfaceless_query)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0