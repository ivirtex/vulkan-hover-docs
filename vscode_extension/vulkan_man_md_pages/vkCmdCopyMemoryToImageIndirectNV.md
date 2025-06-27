# vkCmdCopyMemoryToImageIndirectNV(3) Manual Page

## Name

vkCmdCopyMemoryToImageIndirectNV - Copy data from a memory region into
an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from a memory region to an image object by specifying copy
parameters in a buffer, call:

``` c
// Provided by VK_NV_copy_memory_indirect
void vkCmdCopyMemoryToImageIndirectNV(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             copyBufferAddress,
    uint32_t                                    copyCount,
    uint32_t                                    stride,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    const VkImageSubresourceLayers*             pImageSubresources);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `copyBufferAddress` is the buffer address specifying the copy
  parameters. This buffer is laid out in memory as an array of
  [VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageIndirectCommandNV.html)
  structures.

- `copyCount` is the number of copies to execute, and can be zero.

- `stride` is the byte stride between successive sets of copy
  parameters.

- `dstImage` is the destination image.

- `dstImageLayout` is the layout of the destination image subresources
  for the copy.

- `pImageSubresources` is a pointer to an array of size `copyCount` of
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) used to
  specify the specific image subresource of the destination image data
  for that copy.

## <a href="#_description" class="anchor"></a>Description

Each region in `copyBufferAddress` is copied from the source memory
region to an image region in the destination image. If the destination
image is of type `VK_IMAGE_TYPE_3D`, the starting slice and number of
slices to copy are specified in `pImageSubresources->baseArrayLayer` and
`pImageSubresources->layerCount` respectively. The copy **must** be
performed on a queue that supports indirect copy operations, see
[VkPhysicalDeviceCopyMemoryIndirectPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesNV.html).

Valid Usage

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-None-07660"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-None-07660"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-None-07660  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-indirectCopy"
  target="_blank" rel="noopener"><code>indirectCopy</code></a> feature
  **must** be enabled

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07661"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07661"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07661  
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-aspectMask-07662"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-aspectMask-07662"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-aspectMask-07662  
  The `aspectMask` member for every subresource in `pImageSubresources`
  **must** only have a single bit set

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07663"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07663"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07663  
  The image region specified by each element in `copyBufferAddress`
  **must** be a region that is contained within `dstImage`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07664"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07664"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07664  
  `dstImage` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07665"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07665"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07665  
  If `dstImage` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

<!-- -->

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07973"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07973"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07973  
  `dstImage` **must** have a sample count equal to
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07667"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07667"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07667  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` at the time this command is executed on a `VkDevice`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07669"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07669"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07669  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`,
  `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-mipLevel-07670"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-mipLevel-07670"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-mipLevel-07670  
  The specified `mipLevel` of each region **must** be less than the
  `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-layerCount-08764"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-layerCount-08764"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-layerCount-08764  
  If `layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, the specified
  `baseArrayLayer` + `layerCount` of each region **must** be less than
  or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-imageOffset-07672"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-imageOffset-07672"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-imageOffset-07672  
  The `imageOffset` and `imageExtent` members of each region **must**
  respect the image transfer granularity requirements of
  `commandBuffer`’s command pool’s queue family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07673"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07673"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07673  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-07674"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-07674"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-07674  
  If the queue family used to create the
  [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) which `commandBuffer` was
  allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each
  region, the `aspectMask` member of `pImageSubresources` **must** not
  be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-imageOffset-07675"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-imageOffset-07675"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-imageOffset-07675  
  For each region in `copyBufferAddress`, `imageOffset.y` and
  (`imageExtent.height` + `imageOffset.y`) **must** both be greater than
  or equal to `0` and less than or equal to the height of the specified
  subresource

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-offset-07676"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-offset-07676"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-offset-07676  
  `offset` **must** be 4 byte aligned

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-stride-07677"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-stride-07677"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-stride-07677  
  `stride` **must** be a multiple of `4` and **must** be greater than or
  equal to sizeof(`VkCopyMemoryToImageIndirectCommandNV`)

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-parameter"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-parameter"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-parameter"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-parameter"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a
  href="#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-parameter"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-parameter"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a
  href="#VUID-vkCmdCopyMemoryToImageIndirectNV-pImageSubresources-parameter"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-pImageSubresources-parameter"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-pImageSubresources-parameter  
  `pImageSubresources` **must** be a valid pointer to an array of
  `copyCount` valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structures

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-recording"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-recording"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-renderpass"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-renderpass"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-videocoding"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-videocoding"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-copyCount-arraylength"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-copyCount-arraylength"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-copyCount-arraylength  
  `copyCount` **must** be greater than `0`

- <a href="#VUID-vkCmdCopyMemoryToImageIndirectNV-commonparent"
  id="VUID-vkCmdCopyMemoryToImageIndirectNV-commonparent"></a>
  VUID-vkCmdCopyMemoryToImageIndirectNV-commonparent  
  Both of `commandBuffer`, and `dstImage` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_copy_memory_indirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_copy_memory_indirect.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyMemoryToImageIndirectNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
