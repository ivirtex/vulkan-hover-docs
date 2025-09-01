# VkDebugUtilsLabelEXT(3) Manual Page

## Name

VkDebugUtilsLabelEXT - Specify parameters of a label region



## [](#_c_specification)C Specification

The `VkDebugUtilsLabelEXT` structure is defined as:

```c++
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsLabelEXT {
    VkStructureType    sType;
    const void*        pNext;
    const char*        pLabelName;
    float              color[4];
} VkDebugUtilsLabelEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pLabelName` is a pointer to a null-terminated UTF-8 string containing the name of the label.
- `color` is an optional RGBA color value that can be associated with the label. A particular implementation **may** choose to ignore this color value. The values contain RGBA values in order, in the range 0.0 to 1.0. If all elements in `color` are 0.0, then it is ignored.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDebugUtilsLabelEXT-sType-sType)VUID-VkDebugUtilsLabelEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT`
- [](#VUID-VkDebugUtilsLabelEXT-pNext-pNext)VUID-VkDebugUtilsLabelEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDebugUtilsLabelEXT-pLabelName-parameter)VUID-VkDebugUtilsLabelEXT-pLabelName-parameter  
  `pLabelName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginDebugUtilsLabelEXT.html), [vkCmdInsertDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdInsertDebugUtilsLabelEXT.html), [vkQueueBeginDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBeginDebugUtilsLabelEXT.html), [vkQueueInsertDebugUtilsLabelEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueInsertDebugUtilsLabelEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugUtilsLabelEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0