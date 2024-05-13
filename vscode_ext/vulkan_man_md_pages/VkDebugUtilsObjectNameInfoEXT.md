# VkDebugUtilsObjectNameInfoEXT(3) Manual Page

## Name

VkDebugUtilsObjectNameInfoEXT - Specify parameters of a name to give to
an object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDebugUtilsObjectNameInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsObjectNameInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkObjectType       objectType;
    uint64_t           objectHandle;
    const char*        pObjectName;
} VkDebugUtilsObjectNameInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html) specifying the
  type of the object to be named.

- `objectHandle` is the object to be named.

- `pObjectName` is either `NULL` or a null-terminated UTF-8 string
  specifying the name to apply to `objectHandle`.

## <a href="#_description" class="anchor"></a>Description

Applications **may** change the name associated with an object simply by
calling `vkSetDebugUtilsObjectNameEXT` again with a new string. If
`pObjectName` is either `NULL` or an empty string, then any previously
set name is removed.

The <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-graphicsPipelineLibrary"
target="_blank" rel="noopener"><code>graphicsPipelineLibrary</code></a>
feature allows the specification of pipelines without the creation of
[VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) objects beforehand. In order to
continue to allow naming these shaders independently,
`VkDebugUtilsObjectNameInfoEXT` **can** be included in the `pNext` chain
of
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
which associates a static name with that particular shader.

Valid Usage

- <a href="#VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02589"
  id="VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02589"></a>
  VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02589  
  If `objectType` is `VK_OBJECT_TYPE_UNKNOWN`, `objectHandle` **must**
  not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02590"
  id="VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02590"></a>
  VUID-VkDebugUtilsObjectNameInfoEXT-objectType-02590  
  If `objectType` is not `VK_OBJECT_TYPE_UNKNOWN`, `objectHandle`
  **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a valid Vulkan
  handle of the type associated with `objectType` as defined in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debugging-object-types"
  target="_blank" rel="noopener"><code>VkObjectType</code> and Vulkan
  Handle Relationship</a> table

Valid Usage (Implicit)

- <a href="#VUID-VkDebugUtilsObjectNameInfoEXT-sType-sType"
  id="VUID-VkDebugUtilsObjectNameInfoEXT-sType-sType"></a>
  VUID-VkDebugUtilsObjectNameInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT`

- <a href="#VUID-VkDebugUtilsObjectNameInfoEXT-objectType-parameter"
  id="VUID-VkDebugUtilsObjectNameInfoEXT-objectType-parameter"></a>
  VUID-VkDebugUtilsObjectNameInfoEXT-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html)
  value

- <a href="#VUID-VkDebugUtilsObjectNameInfoEXT-pObjectName-parameter"
  id="VUID-VkDebugUtilsObjectNameInfoEXT-pObjectName-parameter"></a>
  VUID-VkDebugUtilsObjectNameInfoEXT-pObjectName-parameter  
  If `pObjectName` is not `NULL`, `pObjectName` **must** be a
  null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataEXT.html),
[VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSetDebugUtilsObjectNameEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetDebugUtilsObjectNameEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugUtilsObjectNameInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
