# VK\_AMD\_negative\_viewport\_height(3) Manual Page

## Name

VK\_AMD\_negative\_viewport\_height - device extension



## [](#_registered_extension_number)Registered Extension Number

36

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Obsoleted* by [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Matthaeus G. Chajdas [\[GitHub\]anteru](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_negative_viewport_height%5D%20%40anteru%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_negative_viewport_height%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-09-02

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Graham Sellers, AMD
- Baldur Karlsson

## [](#_description)Description

This extension allows an application to specify a negative viewport height. The result is that the viewport transformation will flip along the y-axis.

- Allow negative height to be specified in the [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html)::`height` field to perform y-inversion of the clip-space to framebuffer-space transform. This allows apps to avoid having to use `gl_Position.y = -gl_Position.y` in shaders also targeting other APIs.

## [](#_obsoletion_by_vk_khr_maintenance1_and_vulkan_1_1)Obsoletion by `VK_KHR_maintenance1` and Vulkan 1.1

Functionality in this extension is included in the `VK_KHR_maintenance1` extension and subsequently Vulkan 1.1. Due to some slight behavioral differences, this extension **must** not be enabled alongside `VK_KHR_maintenance1`, or in an instance created with version 1.1 or later requested in [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion`.

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_NEGATIVE_VIEWPORT_HEIGHT_EXTENSION_NAME`
- `VK_AMD_NEGATIVE_VIEWPORT_HEIGHT_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 1, 2016-09-02 (Matthaeus Chajdas)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_negative_viewport_height).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0