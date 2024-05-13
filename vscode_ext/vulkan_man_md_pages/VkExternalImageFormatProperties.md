# VkExternalImageFormatProperties(3) Manual Page

## Name

VkExternalImageFormatProperties - Structure specifying supported
external handle properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExternalImageFormatProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExternalImageFormatProperties {
    VkStructureType               sType;
    void*                         pNext;
    VkExternalMemoryProperties    externalMemoryProperties;
} VkExternalImageFormatProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_external_memory_capabilities
typedef VkExternalImageFormatProperties VkExternalImageFormatPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `externalMemoryProperties` is a
  [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html)
  structure specifying various capabilities of the external handle type
  when used with the specified image creation parameters.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExternalImageFormatProperties-sType-sType"
  id="VUID-VkExternalImageFormatProperties-sType-sType"></a>
  VUID-VkExternalImageFormatProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalImageFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
