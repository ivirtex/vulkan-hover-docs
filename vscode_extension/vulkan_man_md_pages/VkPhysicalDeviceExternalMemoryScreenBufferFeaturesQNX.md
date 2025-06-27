# VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX(3) Manual Page

## Name

VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX - Structure
describing QNX Screen Buffer features that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure is
defined as:

``` c
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           screenBufferImport;
} VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX;
```

## <a href="#_members" class="anchor"></a>Members

The members of the
`VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure
describe the following features:

## <a href="#_description" class="anchor"></a>Description

- <span id="features-screenBufferImport"></span> `screenBufferImport`
  indicates whether QNX Screen buffer import functionality is supported.
  If `screenBufferImport` is set to `VK_TRUE`,
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) supports importing
  `_screen_buffer` from applications. In this case, the application is
  responsible for the resource management of the `_screen_buffer`.

|  |  |
|----|----|
| Features | Functionality |
| `screenBufferImport` | [VkImportScreenBufferInfoQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportScreenBufferInfoQNX.html) |
| Always supported<sup>1</sup> | [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html), [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferPropertiesQNX.html), [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html), [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) |

Table 1. Functionality supported for QNX Screen Buffer features
{#features-externalscreenbuffer-table}

1  
Functionality in this row is always available.

The <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalscreenbuffer-table"
target="_blank" rel="noopener">Functionality supported for QNX Screen
buffer features</a> table summarizes the functionality enabled by the
`VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure. Each
entry in the body of the table summarizes the functionality that **can**
be used when the given features are supported and enabled. This
summarizes Valid Usage statements that are added elsewhere in this
specification.

If the `VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX-sType-sType"
  id="VUID-VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX-sType-sType"></a>
  VUID-VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_SCREEN_BUFFER_FEATURES_QNX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_external_memory_screen_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_external_memory_screen_buffer.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
