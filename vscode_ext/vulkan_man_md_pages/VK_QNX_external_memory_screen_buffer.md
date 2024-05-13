# VK_QNX_external_memory_screen_buffer(3) Manual Page

## Name

VK_QNX_external_memory_screen_buffer - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

530

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_sampler_ycbcr_conversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
         and  
         [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)  
         and  
        
[VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Mike Gorchak <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QNX_external_memory_screen_buffer%5D%20@mgorchak-blackberry%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QNX_external_memory_screen_buffer%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mgorchak-blackberry</a>

- Aaron Ruby <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QNX_external_memory_screen_buffer%5D%20@aruby-blackberry%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QNX_external_memory_screen_buffer%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aruby-blackberry</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-05-17

**IP Status**  
No known IP claims.

**Contributors**  
- Mike Gorchak, QNX / Blackberry Limited

- Aaron Ruby, QNX / Blackberry Limited

## <a href="#_description" class="anchor"></a>Description

This extension enables an application to import QNX Screen
`_screen_buffer` objects created outside of the Vulkan device into
Vulkan memory objects, where they can be bound to images and buffers.

Some `_screen_buffer` images have implementation-defined *external
formats* that **may** not correspond to Vulkan formats. Sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion **can** be used to sample from
these images and convert them to a known color space.

`_screen_buffer` is strongly typed, so naming the handle type is
redundant. The internal layout and therefore size of a `_screen_buffer`
image may depend on native usage flags that do not have corresponding
Vulkan counterparts.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferPropertiesQNX.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html):

  - [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkImportScreenBufferInfoQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportScreenBufferInfoQNX.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX.html)

- Extending
  [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferPropertiesQNX.html):

  - [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QNX_EXTERNAL_MEMORY_SCREEN_BUFFER_EXTENSION_NAME`

- `VK_QNX_EXTERNAL_MEMORY_SCREEN_BUFFER_SPEC_VERSION`

- Extending
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html):

  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_QNX`

  - `VK_STRUCTURE_TYPE_IMPORT_SCREEN_BUFFER_INFO_QNX`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_SCREEN_BUFFER_FEATURES_QNX`

  - `VK_STRUCTURE_TYPE_SCREEN_BUFFER_FORMAT_PROPERTIES_QNX`

  - `VK_STRUCTURE_TYPE_SCREEN_BUFFER_PROPERTIES_QNX`

## <a href="#_issues" class="anchor"></a>Issues

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-05-17 (Mike Gorchak)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QNX_external_memory_screen_buffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
