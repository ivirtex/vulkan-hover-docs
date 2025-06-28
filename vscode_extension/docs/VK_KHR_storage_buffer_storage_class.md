# VK\_KHR\_storage\_buffer\_storage\_class(3) Manual Page

## Name

VK\_KHR\_storage\_buffer\_storage\_class - device extension



## [](#_registered_extension_number)Registered Extension Number

132

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_storage\_buffer\_storage\_class](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_storage_buffer_storage_class.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Alexander Galazin [\[GitHub\]alegal-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_storage_buffer_storage_class%5D%20%40alegal-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_storage_buffer_storage_class%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Contributors**

- Alexander Galazin, ARM
- David Neto, Google

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_storage_buffer_storage_class`

This extension provides a new SPIR-V `StorageBuffer` storage class. A `Block`-decorated object in this class is equivalent to a `BufferBlock`-decorated object in the `Uniform` storage class.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_EXTENSION_NAME`
- `VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 1, 2017-03-23 (Alexander Galazin)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_storage_buffer_storage_class)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0