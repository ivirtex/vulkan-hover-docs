# VK_EXT_queue_family_foreign(3) Manual Page

## Name

VK_EXT_queue_family_foreign - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

127

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Lina Versace <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_queue_family_foreign%5D%20@versalinyaa%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_queue_family_foreign%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>versalinyaa</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension defines a special queue family,
`VK_QUEUE_FAMILY_FOREIGN_EXT`, which can be used to transfer ownership
of resources backed by external memory to foreign, external queues. This
is similar to `VK_QUEUE_FAMILY_EXTERNAL_KHR`, defined in
[`VK_KHR_external_memory`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html). The key
differences between the two are:

- The queues represented by `VK_QUEUE_FAMILY_EXTERNAL_KHR` must share
  the same physical device and the same driver version as the current
  [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html). `VK_QUEUE_FAMILY_FOREIGN_EXT` has no
  such restrictions. It can represent devices and drivers from other
  vendors, and can even represent non-Vulkan-capable devices.

- All resources backed by external memory support
  `VK_QUEUE_FAMILY_EXTERNAL_KHR`. Support for
  `VK_QUEUE_FAMILY_FOREIGN_EXT` is more restrictive.

- Applications should expect transitions to/from
  `VK_QUEUE_FAMILY_FOREIGN_EXT` to be more expensive than transitions
  to/from `VK_QUEUE_FAMILY_EXTERNAL_KHR`.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_QUEUE_FAMILY_FOREIGN_EXTENSION_NAME`

- `VK_EXT_QUEUE_FAMILY_FOREIGN_SPEC_VERSION`

- `VK_QUEUE_FAMILY_FOREIGN_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-11-01 (Lina Versace)

  - Squashed internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_queue_family_foreign"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
