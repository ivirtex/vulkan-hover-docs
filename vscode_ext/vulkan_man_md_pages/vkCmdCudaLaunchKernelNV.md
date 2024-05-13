# vkCmdCudaLaunchKernelNV(3) Manual Page

## Name

vkCmdCudaLaunchKernelNV - Dispatch compute work items



## <a href="#_c_specification" class="anchor"></a>C Specification

To record a CUDA kernel launch, call:

``` c
// Provided by VK_NV_cuda_kernel_launch
void vkCmdCudaLaunchKernelNV(
    VkCommandBuffer                             commandBuffer,
    const VkCudaLaunchInfoNV*                   pLaunchInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pLaunchInfo` is a pointer to a
  [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaLaunchInfoNV.html) structure in which the
  grid (similar to workgroup) dimension, function handle and related
  arguments are defined.

## <a href="#_description" class="anchor"></a>Description

When the command is executed, a global workgroup consisting of
`gridDimX` × `gridDimY` × `gridDimZ` local workgroups is assembled.

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCudaLaunchKernelNV-commandBuffer-parameter"
  id="VUID-vkCmdCudaLaunchKernelNV-commandBuffer-parameter"></a>
  VUID-vkCmdCudaLaunchKernelNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCudaLaunchKernelNV-pLaunchInfo-parameter"
  id="VUID-vkCmdCudaLaunchKernelNV-pLaunchInfo-parameter"></a>
  VUID-vkCmdCudaLaunchKernelNV-pLaunchInfo-parameter  
  `pLaunchInfo` **must** be a valid pointer to a valid
  [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaLaunchInfoNV.html) structure

- <a href="#VUID-vkCmdCudaLaunchKernelNV-commandBuffer-recording"
  id="VUID-vkCmdCudaLaunchKernelNV-commandBuffer-recording"></a>
  VUID-vkCmdCudaLaunchKernelNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCudaLaunchKernelNV-commandBuffer-cmdpool"
  id="VUID-vkCmdCudaLaunchKernelNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdCudaLaunchKernelNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdCudaLaunchKernelNV-videocoding"
  id="VUID-vkCmdCudaLaunchKernelNV-videocoding"></a>
  VUID-vkCmdCudaLaunchKernelNV-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

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
<tr class="header">
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
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaLaunchInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCudaLaunchKernelNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
