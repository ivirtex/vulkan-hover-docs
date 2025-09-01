# VK\_EXT\_queue\_family\_foreign(3) Manual Page

## Name

VK\_EXT\_queue\_family\_foreign - device extension



## [](#_registered_extension_number)Registered Extension Number

127

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Lina Versace [\[GitHub\]linyaa-kiwi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_queue_family_foreign%5D%20%40linyaa-kiwi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_queue_family_foreign%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-11-01

**IP Status**

No known IP claims.

**Contributors**

- Lina Versace, Google
- James Jones, NVIDIA
- Faith Ekstrand, Intel
- Jesse Hall, Google
- Daniel Rakos, AMD
- Ray Smith, ARM

## [](#_description)Description

This extension defines a special queue family, `VK_QUEUE_FAMILY_FOREIGN_EXT`, which can be used to transfer ownership of resources backed by external memory to foreign, external queues. This is similar to `VK_QUEUE_FAMILY_EXTERNAL_KHR`, defined in `VK_KHR_external_memory`. The key differences between the two are:

- The queues represented by `VK_QUEUE_FAMILY_EXTERNAL_KHR` must share the same physical device and the same driver version as the current [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html). `VK_QUEUE_FAMILY_FOREIGN_EXT` has no such restrictions. It can represent devices and drivers from other vendors, and can even represent non-Vulkan-capable devices.
- All resources backed by external memory support `VK_QUEUE_FAMILY_EXTERNAL_KHR`. Support for `VK_QUEUE_FAMILY_FOREIGN_EXT` is more restrictive.
- Applications should expect transitions to/from `VK_QUEUE_FAMILY_FOREIGN_EXT` to be more expensive than transitions to/from `VK_QUEUE_FAMILY_EXTERNAL_KHR`.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_QUEUE_FAMILY_FOREIGN_EXTENSION_NAME`
- `VK_EXT_QUEUE_FAMILY_FOREIGN_SPEC_VERSION`
- `VK_QUEUE_FAMILY_FOREIGN_EXT`

## [](#_version_history)Version History

- Revision 1, 2017-11-01 (Lina Versace)
  
  - Squashed internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_queue_family_foreign).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0