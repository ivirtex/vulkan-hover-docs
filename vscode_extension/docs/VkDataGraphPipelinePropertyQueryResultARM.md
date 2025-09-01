# VkDataGraphPipelinePropertyQueryResultARM(3) Manual Page

## Name

VkDataGraphPipelinePropertyQueryResultARM - Structure describing a data graph pipeline property query or result



## [](#_c_specification)C Specification

The `VkDataGraphPipelinePropertyQueryResultARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelinePropertyQueryResultARM {
    VkStructureType                   sType;
    const void*                       pNext;
    VkDataGraphPipelinePropertyARM    property;
    VkBool32                          isText;
    size_t                            dataSize;
    void*                             pData;
} VkDataGraphPipelinePropertyQueryResultARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `property` is a [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html) specifying the property of the data graph pipeline being queried.
- `isText` specifies whether the returned data is text or opaque data. If `isText` is `VK_TRUE` then the data returned in `pData` is text and guaranteed to be a null-terminated UTF-8 string.
- `dataSize` is an integer related to the size, in bytes, of the data, as described below.
- `pData` is either `NULL` or a pointer to a block of memory into which the implementation will return the property data.

## [](#_description)Description

If `pData` is `NULL`, then the size, in bytes, of the property data is returned in `dataSize`. Otherwise, `dataSize` must be the size of the buffer, in bytes, pointed to by `pData` and on return `dataSize` is overwritten with the number of bytes of data actually written to `pData` including any trailing NUL character. If `dataSize` is less than the size, in bytes, of the property data, at most `dataSize` bytes of data will be written to `pData`, and `VK_INCOMPLETE` will be returned by [vkGetDataGraphPipelinePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelinePropertiesARM.html) instead of `VK_SUCCESS`, to indicate that not all the available property data was returned. If `isText` is `VK_TRUE` and `pData` is not `NULL` and `dataSize` is not zero, the last byte written to `pData` will be a NUL character.

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelinePropertyQueryResultARM-sType-sType)VUID-VkDataGraphPipelinePropertyQueryResultARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_PROPERTY_QUERY_RESULT_ARM`
- [](#VUID-VkDataGraphPipelinePropertyQueryResultARM-pNext-pNext)VUID-VkDataGraphPipelinePropertyQueryResultARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelinePropertyQueryResultARM-property-parameter)VUID-VkDataGraphPipelinePropertyQueryResultARM-property-parameter  
  `property` **must** be a valid [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html) value
- [](#VUID-VkDataGraphPipelinePropertyQueryResultARM-pData-parameter)VUID-VkDataGraphPipelinePropertyQueryResultARM-pData-parameter  
  If `dataSize` is not `0`, and `pData` is not `NULL`, `pData` **must** be a valid pointer to an array of `dataSize` bytes

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDataGraphPipelinePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelinePropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelinePropertyQueryResultARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0