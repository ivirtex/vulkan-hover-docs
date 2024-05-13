# VK_EXT_private_data(3) Manual Page

## Name

VK_EXT_private_data - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

296

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Rusch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_private_data%5D%20@mattruschnv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_private_data%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mattruschnv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-03-25

**IP Status**  
No known IP claims.

**Contributors**  
- Matthew Rusch, NVIDIA

- Nuno Subtil, NVIDIA

- Piers Daniell, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension is a device extension which enables attaching arbitrary
payloads to Vulkan objects. It introduces the idea of private data slots
as a means of storing a 64-bit unsigned integer of application defined
data. Private data slots can be created or destroyed any time an
associated device is available. Private data slots can be reserved at
device creation time, and limiting use to the amount reserved will allow
the extension to exhibit better performance characteristics.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkPrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotEXT.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreatePrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePrivateDataSlotEXT.html)

- [vkDestroyPrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPrivateDataSlotEXT.html)

- [vkGetPrivateDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPrivateDataEXT.html)

- [vkSetPrivateDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetPrivateDataEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPrivateDataSlotCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotCreateInfoEXT.html)

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkDevicePrivateDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevicePrivateDataCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePrivateDataFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePrivateDataFeaturesEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPrivateDataSlotCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PRIVATE_DATA_EXTENSION_NAME`

- `VK_EXT_PRIVATE_DATA_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIVATE_DATA_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
EXT suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_examples" class="anchor"></a>Examples

- In progress

## <a href="#_issues" class="anchor"></a>Issues

\(1\) If I have to create a [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html)
to store and retrieve data on an object, how does this extension help
me? Will I not need to store the
[VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) mapping with each object,
and if I am doing that, I might as well just store the original data!

**RESOLVED:** The [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) can be
thought of as an opaque index into storage that is reserved in each
object. That is, you can use the same
[VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) with each object for a
specific piece of information. For example, if a layer wishes to track
per-object information, the layer only needs to allocate one
[VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlot.html) per device and it can use
that private data slot for all of the device’s child objects. This
allows multiple layers to store private data without conflicting with
each other’s and/or the application’s private data.

\(2\) What if I need to store more than 64-bits of information per
object?

**RESOLVED:** The data that you store per object could be a pointer to
another object or structure of your own allocation.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-01-15 (Matthew Rusch)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_private_data"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
