# VkVideoInlineQueryInfoKHR(3) Manual Page

## Name

VkVideoInlineQueryInfoKHR - Structure specifying inline query
information for video coding commands



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoInlineQueryInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_maintenance1
typedef struct VkVideoInlineQueryInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkQueryPool        queryPool;
    uint32_t           firstQuery;
    uint32_t           queryCount;
} VkVideoInlineQueryInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queryPool` is `VK_NULL_HANDLE` or a valid handle to a
  [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) object that will manage the results of
  the queries.

- `firstQuery` is the query index within the query pool that will
  contain the query results for the first video coding operation. The
  query results of subsequent video coding operations will be contained
  by subsequent query indices.

- `queryCount` is the number of queries to execute.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>In practice, if <code>queryPool</code> is not
  <code>VK_NULL_HANDLE</code>, then <code>queryCount</code> will always
  have to match the number of video coding operations issued by the video
  coding command this structure is specified to, meaning that using inline
  queries in a video coding command will always execute a query for each
  issued video coding operation.</p></td>
  </tr>
  </tbody>
  </table>

## <a href="#_description" class="anchor"></a>Description

This structure **can** be included in the `pNext` chain of the input
parameter structure of video coding commands.

- In the `pNext` chain of the `pDecodeInfo` parameter of the
  [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html) command to execute a
  query for each video decode operation issued by the command.

- In the `pNext` chain of the `pEncodeInfo` parameter of the
  [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html) command to execute a
  query for each video encode operation issued by the command.

Valid Usage

- <a href="#VUID-VkVideoInlineQueryInfoKHR-queryPool-08372"
  id="VUID-VkVideoInlineQueryInfoKHR-queryPool-08372"></a>
  VUID-VkVideoInlineQueryInfoKHR-queryPool-08372  
  If `queryPool` is not `VK_NULL_HANDLE`, then `firstQuery` **must** be
  less than the number of queries in `queryPool`

- <a href="#VUID-VkVideoInlineQueryInfoKHR-queryPool-08373"
  id="VUID-VkVideoInlineQueryInfoKHR-queryPool-08373"></a>
  VUID-VkVideoInlineQueryInfoKHR-queryPool-08373  
  If `queryPool` is not `VK_NULL_HANDLE`, then the sum of `firstQuery`
  and `queryCount` **must** be less than or equal to the number of
  queries in `queryPool`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoInlineQueryInfoKHR-sType-sType"
  id="VUID-VkVideoInlineQueryInfoKHR-sType-sType"></a>
  VUID-VkVideoInlineQueryInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_INLINE_QUERY_INFO_KHR`

- <a href="#VUID-VkVideoInlineQueryInfoKHR-queryPool-parameter"
  id="VUID-VkVideoInlineQueryInfoKHR-queryPool-parameter"></a>
  VUID-VkVideoInlineQueryInfoKHR-queryPool-parameter  
  If `queryPool` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_maintenance1.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoInlineQueryInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
