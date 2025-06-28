# VK\_KHR\_surface\_protected\_capabilities(3) Manual Page

## Name

VK\_KHR\_surface\_protected\_capabilities - instance extension



## [](#_registered_extension_number)Registered Extension Number

240

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)

## [](#_contact)Contact

- Sandeep Shinde [\[GitHub\]sashinde](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface_protected_capabilities%5D%20%40sashinde%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface_protected_capabilities%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-12-18

**IP Status**

No known IP claims.

**Contributors**

- Sandeep Shinde, NVIDIA
- James Jones, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension extends [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html), providing applications a way to query whether swapchains **can** be created with the `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` flag set.

Vulkan 1.1 added (optional) support for protect memory and protected resources including buffers (`VK_BUFFER_CREATE_PROTECTED_BIT`), images (`VK_IMAGE_CREATE_PROTECTED_BIT`), and swapchains (`VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR`). However, on implementations which support multiple windowing systems, not all window systems **may** be able to provide a protected display path.

This extension provides a way to query if a protected swapchain created for a surface (and thus a specific windowing system) **can** be displayed on screen. It extends the existing [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html) structure with a new [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceProtectedCapabilitiesKHR.html) structure from which the application **can** obtain information about support for protected swapchain creation through [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).

## [](#_new_structures)New Structures

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceProtectedCapabilitiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SURFACE_PROTECTED_CAPABILITIES_EXTENSION_NAME`
- `VK_KHR_SURFACE_PROTECTED_CAPABILITIES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-12-18 (Sandeep Shinde, Daniel Koch)
  
  - Internal revisions.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_surface_protected_capabilities)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0