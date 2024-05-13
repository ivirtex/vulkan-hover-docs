# VkDebugUtilsLabelEXT(3) Manual Page

## Name

VkDebugUtilsLabelEXT - Specify parameters of a label region



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDebugUtilsLabelEXT` structure is defined as:

``` c
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsLabelEXT {
    VkStructureType    sType;
    const void*        pNext;
    const char*        pLabelName;
    float              color[4];
} VkDebugUtilsLabelEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pLabelName` is a pointer to a null-terminated UTF-8 string containing
  the name of the label.

- `color` is an optional RGBA color value that can be associated with
  the label. A particular implementation **may** choose to ignore this
  color value. The values contain RGBA values in order, in the range 0.0
  to 1.0. If all elements in `color` are set to 0.0 then it is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDebugUtilsLabelEXT-sType-sType"
  id="VUID-VkDebugUtilsLabelEXT-sType-sType"></a>
  VUID-VkDebugUtilsLabelEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT`

- <a href="#VUID-VkDebugUtilsLabelEXT-pNext-pNext"
  id="VUID-VkDebugUtilsLabelEXT-pNext-pNext"></a>
  VUID-VkDebugUtilsLabelEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDebugUtilsLabelEXT-pLabelName-parameter"
  id="VUID-VkDebugUtilsLabelEXT-pLabelName-parameter"></a>
  VUID-VkDebugUtilsLabelEXT-pLabelName-parameter  
  `pLabelName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginDebugUtilsLabelEXT.html),
[vkCmdInsertDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdInsertDebugUtilsLabelEXT.html),
[vkQueueBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueBeginDebugUtilsLabelEXT.html),
[vkQueueInsertDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueInsertDebugUtilsLabelEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsLabelEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
