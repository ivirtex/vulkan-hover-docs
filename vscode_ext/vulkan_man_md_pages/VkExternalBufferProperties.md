# VkExternalBufferProperties(3) Manual Page

## Name

VkExternalBufferProperties - Structure specifying supported external
handle capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExternalBufferProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExternalBufferProperties {
    VkStructureType               sType;
    void*                         pNext;
    VkExternalMemoryProperties    externalMemoryProperties;
} VkExternalBufferProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_external_memory_capabilities
typedef VkExternalBufferProperties VkExternalBufferPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `externalMemoryProperties` is a
  [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html)
  structure specifying various capabilities of the external handle type
  when used with the specified buffer creation parameters.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExternalBufferProperties-sType-sType"
  id="VUID-VkExternalBufferProperties-sType-sType"></a>
  VUID-VkExternalBufferProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES`

- <a href="#VUID-VkExternalBufferProperties-pNext-pNext"
  id="VUID-VkExternalBufferProperties-pNext-pNext"></a>
  VUID-VkExternalBufferProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html),
[vkGetPhysicalDeviceExternalBufferPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalBufferProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
