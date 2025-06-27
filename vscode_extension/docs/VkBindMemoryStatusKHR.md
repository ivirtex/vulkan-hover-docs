# VkBindMemoryStatusKHR(3) Manual Page

## Name

VkBindMemoryStatusKHR - Structure specifying where to return memory
binding status



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindMemoryStatusKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6
typedef struct VkBindMemoryStatusKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkResult*          pResult;
} VkBindMemoryStatusKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pResult` is a pointer to a `VkResult` value.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) or
[VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) includes a
`VkBindMemoryStatusKHR` structure, then the
`VkBindMemoryStatusKHR`::`pResult` will be populated with a value
describing the result of the corresponding memory binding operation.

Valid Usage (Implicit)

- <a href="#VUID-VkBindMemoryStatusKHR-sType-sType"
  id="VUID-VkBindMemoryStatusKHR-sType-sType"></a>
  VUID-VkBindMemoryStatusKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_MEMORY_STATUS_KHR`

- <a href="#VUID-VkBindMemoryStatusKHR-pResult-parameter"
  id="VUID-VkBindMemoryStatusKHR-pResult-parameter"></a>
  VUID-VkBindMemoryStatusKHR-pResult-parameter  
  `pResult` **must** be a valid pointer to a [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindMemoryStatusKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
