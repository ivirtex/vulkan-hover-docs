# VkDebugUtilsObjectNameInfoEXT(3) Manual Page

## Name

VkDebugUtilsObjectNameInfoEXT - Specify parameters of a name to give to an object



## [](#_c_specification)C Specification

The `VkDebugUtilsObjectNameInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsObjectNameInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkObjectType       objectType;
    uint64_t           objectHandle;
    const char*        pObjectName;
} VkDebugUtilsObjectNameInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) specifying the type of the object to be named.
- `objectHandle` is the object to be named.
- `pObjectName` is either `NULL` or a null-terminated UTF-8 string specifying the name to apply to `objectHandle`.

## [](#_description)Description

Applications **may** change the name associated with an object simply by calling `vkSetDebugUtilsObjectNameEXT` again with a new string. If `pObjectName` is either `NULL` or an empty string, then any previously set name is removed.

The [`graphicsPipelineLibrary`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-graphicsPipelineLibrary) feature allows the specification of pipelines without the creation of [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) objects beforehand. In order to continue to allow naming these shaders independently, `VkDebugUtilsObjectNameInfoEXT` **can** be included in the `pNext` chain of [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), which associates a static name with that particular shader.

Valid Usage

- [](#VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02589)VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02589  
  If `objectType` is `VK_OBJECT_TYPE_UNKNOWN`, `objectHandle` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02590)VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02590  
  If `objectType` is not `VK_OBJECT_TYPE_UNKNOWN`, `objectHandle` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a valid Vulkan handle of the type associated with `objectType` as defined in the [`VkObjectType` and Vulkan Handle Relationship](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debugging-object-types) table

Valid Usage (Implicit)

- [](#VUID-VkDebugUtilsObjectNameInfoEXT-sType-sType)VUID-VkDebugUtilsObjectNameInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT`
- [](#VUID-VkDebugUtilsObjectNameInfoEXT-objectType-parameter)VUID-VkDebugUtilsObjectNameInfoEXT-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) value
- [](#VUID-VkDebugUtilsObjectNameInfoEXT-pObjectName-parameter)VUID-VkDebugUtilsObjectNameInfoEXT-pObjectName-parameter  
  If `pObjectName` is not `NULL`, `pObjectName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html), [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkSetDebugUtilsObjectNameEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetDebugUtilsObjectNameEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugUtilsObjectNameInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0