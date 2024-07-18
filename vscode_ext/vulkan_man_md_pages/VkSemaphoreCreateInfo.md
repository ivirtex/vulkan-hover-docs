# VkSemaphoreCreateInfo(3) Manual Page

## Name

VkSemaphoreCreateInfo - Structure specifying parameters of a newly
created semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSemaphoreCreateInfo {
    VkStructureType           sType;
    const void*               pNext;
    VkSemaphoreCreateFlags    flags;
} VkSemaphoreCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSemaphoreCreateInfo-pNext-06789"
  id="VUID-VkSemaphoreCreateInfo-pNext-06789"></a>
  VUID-VkSemaphoreCreateInfo-pNext-06789  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreCreateInfo-sType-sType"
  id="VUID-VkSemaphoreCreateInfo-sType-sType"></a>
  VUID-VkSemaphoreCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO`

- <a href="#VUID-VkSemaphoreCreateInfo-pNext-pNext"
  id="VUID-VkSemaphoreCreateInfo-pNext-pNext"></a>
  VUID-VkSemaphoreCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
  [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html),
  [VkExportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreWin32HandleInfoKHR.html),
  [VkImportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalSharedEventInfoEXT.html),
  [VkQueryLowLatencySupportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryLowLatencySupportNV.html), or
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)

- <a href="#VUID-VkSemaphoreCreateInfo-sType-unique"
  id="VUID-VkSemaphoreCreateInfo-sType-unique"></a>
  VUID-VkSemaphoreCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

- <a href="#VUID-VkSemaphoreCreateInfo-flags-zerobitmask"
  id="VUID-VkSemaphoreCreateInfo-flags-zerobitmask"></a>
  VUID-VkSemaphoreCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSemaphoreCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
