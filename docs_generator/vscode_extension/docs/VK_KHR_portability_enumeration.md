# VK\_KHR\_portability\_enumeration(3) Manual Page

## Name

VK\_KHR\_portability\_enumeration - instance extension



## [](#_registered_extension_number)Registered Extension Number

395

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Charles Giessen [\[GitHub\]charles-lunarg](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_portability_enumeration%5D%20%40charles-lunarg%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_portability_enumeration%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-02

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Interacts with `VK_KHR_portability_subset`

**Contributors**

- Lenny Komow, LunarG
- Charles Giessen, LunarG

## [](#_description)Description

This extension allows applications to control whether devices that expose the `VK_KHR_portability_subset` extension are included in the results of physical device enumeration. Since devices which support the `VK_KHR_portability_subset` extension are not fully conformant Vulkan implementations, the Vulkan loader does not report those devices unless the application explicitly asks for them. This prevents applications which may not be aware of non-conformant devices from accidentally using them, as any device which supports the `VK_KHR_portability_subset` extension mandates that the extension must be enabled if that device is used.

This extension is implemented in the loader.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME`
- `VK_KHR_PORTABILITY_ENUMERATION_SPEC_VERSION`
- Extending [VkInstanceCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateFlagBits.html):
  
  - `VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2021-06-02 (Lenny Komow)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_portability_enumeration)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0