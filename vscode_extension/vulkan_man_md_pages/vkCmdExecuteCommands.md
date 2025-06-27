# vkCmdExecuteCommands(3) Manual Page

## Name

vkCmdExecuteCommands - Execute a secondary command buffer from a primary
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Secondary command buffers **must** not be directly submitted to a queue.
To record a secondary command buffer to execute as part of a primary
command buffer, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdExecuteCommands(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    commandBufferCount,
    const VkCommandBuffer*                      pCommandBuffers);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is a handle to a primary command buffer that the
  secondary command buffers are executed in.

- `commandBufferCount` is the length of the `pCommandBuffers` array.

- `pCommandBuffers` is a pointer to an array of `commandBufferCount`
  secondary command buffer handles, which are recorded to execute in the
  primary command buffer in the order they are listed in the array.

## <a href="#_description" class="anchor"></a>Description

If any element of `pCommandBuffers` was not recorded with the
`VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` flag, and it was recorded
into any other primary command buffer which is currently in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">executable or recording state</a>, that
primary command buffer becomes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid</a>.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
feature is enabled it is valid usage for `vkCmdExecuteCommands` to also
be recorded to a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
target="_blank" rel="noopener">secondary command buffer</a>.

Valid Usage

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00088"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00088"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00088  
  Each element of `pCommandBuffers` **must** have been allocated with a
  `level` of `VK_COMMAND_BUFFER_LEVEL_SECONDARY`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00089"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00089"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00089  
  Each element of `pCommandBuffers` **must** be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending or executable state</a>

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00091"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00091"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00091  
  If any element of `pCommandBuffers` was not recorded with the
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` flag, it **must** not
  be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00092"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00092"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00092  
  If any element of `pCommandBuffers` was not recorded with the
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` flag, it **must** not
  have already been recorded to `commandBuffer`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00093"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00093"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00093  
  If any element of `pCommandBuffers` was not recorded with the
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` flag, it **must** not
  appear more than once in `pCommandBuffers`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00094"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00094"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00094  
  Each element of `pCommandBuffers` **must** have been allocated from a
  `VkCommandPool` that was created for the same queue family as the
  `VkCommandPool` from which `commandBuffer` was allocated

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00096"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00096"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00096  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance, each element of `pCommandBuffers` **must** have been
  recorded with the `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00099"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00099"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00099  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance, and any element of `pCommandBuffers` was recorded with
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)::`framebuffer`
  not equal to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), that
  `VkFramebuffer` **must** match the `VkFramebuffer` used in the current
  render pass instance

- <a href="#VUID-vkCmdExecuteCommands-contents-09680"
  id="VUID-vkCmdExecuteCommands-contents-09680"></a>
  VUID-vkCmdExecuteCommands-contents-09680  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html),
  and [vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html) has not been called in
  the current render pass instance, the `contents` parameter of
  [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html) **must** have been
  set to `VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS` , or
  `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT`

- <a href="#VUID-vkCmdExecuteCommands-None-09681"
  id="VUID-vkCmdExecuteCommands-None-09681"></a>
  VUID-vkCmdExecuteCommands-None-09681  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html),
  and [vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html) has been called in the
  current render pass instance, the `contents` parameter of the last
  call to [vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html) **must** have been
  set to `VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS` , or
  `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-06019"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-06019"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-06019  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html),
  each element of `pCommandBuffers` **must** have been recorded with
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)::`subpass`
  set to the index of the subpass which the given command buffer will be
  executed in

- <a href="#VUID-vkCmdExecuteCommands-pBeginInfo-06020"
  id="VUID-vkCmdExecuteCommands-pBeginInfo-06020"></a>
  VUID-vkCmdExecuteCommands-pBeginInfo-06020  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html),
  the render passes specified in the
  `pBeginInfo->pInheritanceInfo->renderPass` members of the
  [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html) commands used to
  begin recording each element of `pCommandBuffers` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">compatible</a> with the current render
  pass

- <a href="#VUID-vkCmdExecuteCommands-pNext-02865"
  id="VUID-vkCmdExecuteCommands-pNext-02865"></a>
  VUID-vkCmdExecuteCommands-pNext-02865  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance that included
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)
  in the `pNext` chain of
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html), then each element
  of `pCommandBuffers` **must** have been recorded with
  [VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html)
  in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)

- <a href="#VUID-vkCmdExecuteCommands-pNext-02866"
  id="VUID-vkCmdExecuteCommands-pNext-02866"></a>
  VUID-vkCmdExecuteCommands-pNext-02866  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance that included
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)
  in the `pNext` chain of
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html), then each element
  of `pCommandBuffers` **must** have been recorded with
  [VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html)::`transform`
  identical to
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform`

- <a href="#VUID-vkCmdExecuteCommands-pNext-02867"
  id="VUID-vkCmdExecuteCommands-pNext-02867"></a>
  VUID-vkCmdExecuteCommands-pNext-02867  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance that included
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)
  in the `pNext` chain of
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html), then each element
  of `pCommandBuffers` **must** have been recorded with
  [VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html)::`renderArea`
  identical to
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderArea`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00100"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00100"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00100  
  If `vkCmdExecuteCommands` is not being called within a render pass
  instance, each element of `pCommandBuffers` **must** not have been
  recorded with the `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-00101"
  id="VUID-vkCmdExecuteCommands-commandBuffer-00101"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-00101  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-inheritedQueries"
  target="_blank" rel="noopener"><code>inheritedQueries</code></a>
  feature is not enabled, `commandBuffer` **must** not have any queries
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-00102"
  id="VUID-vkCmdExecuteCommands-commandBuffer-00102"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-00102  
  If `commandBuffer` has a `VK_QUERY_TYPE_OCCLUSION` query <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>, then each element of
  `pCommandBuffers` **must** have been recorded with
  `VkCommandBufferInheritanceInfo`::`occlusionQueryEnable` set to
  `VK_TRUE`

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-00103"
  id="VUID-vkCmdExecuteCommands-commandBuffer-00103"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-00103  
  If `commandBuffer` has a `VK_QUERY_TYPE_OCCLUSION` query <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>, then each element of
  `pCommandBuffers` **must** have been recorded with
  `VkCommandBufferInheritanceInfo`::`queryFlags` having all bits set
  that are set for the query

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-00104"
  id="VUID-vkCmdExecuteCommands-commandBuffer-00104"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-00104  
  If `commandBuffer` has a `VK_QUERY_TYPE_PIPELINE_STATISTICS` query <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>, then each element of
  `pCommandBuffers` **must** have been recorded with
  `VkCommandBufferInheritanceInfo`::`pipelineStatistics` having all bits
  set that are set in the `VkQueryPool` the query uses

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-00105"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-00105"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-00105  
  Each element of `pCommandBuffers` **must** not begin any query types
  that are <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> in `commandBuffer`

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-07594"
  id="VUID-vkCmdExecuteCommands-commandBuffer-07594"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-07594  
  `commandBuffer` **must** not have any queries other than
  `VK_QUERY_TYPE_OCCLUSION` and `VK_QUERY_TYPE_PIPELINE_STATISTICS` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a>

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-01820"
  id="VUID-vkCmdExecuteCommands-commandBuffer-01820"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-01820  
  If `commandBuffer` is a protected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, each element of `pCommandBuffers` **must** be a
  protected command buffer

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-01821"
  id="VUID-vkCmdExecuteCommands-commandBuffer-01821"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-01821  
  If `commandBuffer` is an unprotected command buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, each element of `pCommandBuffers` **must** be an
  unprotected command buffer

- <a href="#VUID-vkCmdExecuteCommands-None-02286"
  id="VUID-vkCmdExecuteCommands-None-02286"></a>
  VUID-vkCmdExecuteCommands-None-02286  
  This command **must** not be recorded when transform feedback is
  active

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-06533"
  id="VUID-vkCmdExecuteCommands-commandBuffer-06533"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-06533  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance and any recorded command in `commandBuffer` in the current
  subpass will write to an image subresource as an attachment, commands
  recorded in elements of `pCommandBuffers` **must** not read from the
  memory backing that image subresource in any other way

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-06534"
  id="VUID-vkCmdExecuteCommands-commandBuffer-06534"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-06534  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance and any recorded command in `commandBuffer` in the current
  subpass will read from an image subresource used as an attachment in
  any way other than as an attachment, commands recorded in elements of
  `pCommandBuffers` **must** not write to that image subresource as an
  attachment

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-06535"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-06535"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-06535  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance and any recorded command in a given element of
  `pCommandBuffers` will write to an image subresource as an attachment,
  commands recorded in elements of `pCommandBuffers` at a higher index
  **must** not read from the memory backing that image subresource in
  any other way

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-06536"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-06536"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-06536  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance and any recorded command in a given element of
  `pCommandBuffers` will read from an image subresource used as an
  attachment in any way other than as an attachment, commands recorded
  in elements of `pCommandBuffers` at a higher index **must** not write
  to that image subresource as an attachment

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-06021"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-06021"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-06021  
  If `pCommandBuffers` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  there **must** be no action or synchronization commands between that
  render pass instance and any render pass instance that resumes it

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-06022"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-06022"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-06022  
  If `pCommandBuffers` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  there **must** be no render pass instances between that render pass
  instance and any render pass instance that resumes it

- <a href="#VUID-vkCmdExecuteCommands-variableSampleLocations-06023"
  id="VUID-vkCmdExecuteCommands-variableSampleLocations-06023"></a>
  VUID-vkCmdExecuteCommands-variableSampleLocations-06023  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-variableSampleLocations"
  target="_blank" rel="noopener"><code>variableSampleLocations</code></a>
  limit is not supported, and any element of `pCommandBuffers` contains
  any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  where a graphics pipeline has been bound, any pipelines bound in the
  render pass instance that resumes it, or any subsequent render pass
  instances that resume from that one and so on, **must** use the same
  sample locations

- <a href="#VUID-vkCmdExecuteCommands-flags-06024"
  id="VUID-vkCmdExecuteCommands-flags-06024"></a>
  VUID-vkCmdExecuteCommands-flags-06024  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  its [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`flags` parameter
  **must** have included
  `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT`

- <a href="#VUID-vkCmdExecuteCommands-pBeginInfo-06025"
  id="VUID-vkCmdExecuteCommands-pBeginInfo-06025"></a>
  VUID-vkCmdExecuteCommands-pBeginInfo-06025  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  the render passes specified in the
  `pBeginInfo->pInheritanceInfo->renderPass` members of the
  [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html) commands used to
  begin recording each element of `pCommandBuffers` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdExecuteCommands-flags-06026"
  id="VUID-vkCmdExecuteCommands-flags-06026"></a>
  VUID-vkCmdExecuteCommands-flags-06026  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  the `flags` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`flags`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  excluding `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT`

- <a href="#VUID-vkCmdExecuteCommands-colorAttachmentCount-06027"
  id="VUID-vkCmdExecuteCommands-colorAttachmentCount-06027"></a>
  VUID-vkCmdExecuteCommands-colorAttachmentCount-06027  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  the `colorAttachmentCount` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`colorAttachmentCount`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdExecuteCommands-imageView-06028"
  id="VUID-vkCmdExecuteCommands-imageView-06028"></a>
  VUID-vkCmdExecuteCommands-imageView-06028  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  if the `imageView` member of an element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` parameter
  to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the corresponding element of
  the `pColorAttachmentFormats` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the format used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-imageView-07606"
  id="VUID-vkCmdExecuteCommands-imageView-07606"></a>
  VUID-vkCmdExecuteCommands-imageView-07606  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  if the `imageView` member of an element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` parameter
  to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the corresponding element of
  the `pColorAttachmentFormats` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdExecuteCommands-pDepthAttachment-06029"
  id="VUID-vkCmdExecuteCommands-pDepthAttachment-06029"></a>
  VUID-vkCmdExecuteCommands-pDepthAttachment-06029  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  if the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `depthAttachmentFormat` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the format used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pStencilAttachment-06030"
  id="VUID-vkCmdExecuteCommands-pStencilAttachment-06030"></a>
  VUID-vkCmdExecuteCommands-pStencilAttachment-06030  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  if the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `stencilAttachmentFormat` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the format used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pDepthAttachment-06774"
  id="VUID-vkCmdExecuteCommands-pDepthAttachment-06774"></a>
  VUID-vkCmdExecuteCommands-pDepthAttachment-06774  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) was
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `depthAttachmentFormat` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdExecuteCommands-pStencilAttachment-06775"
  id="VUID-vkCmdExecuteCommands-pStencilAttachment-06775"></a>
  VUID-vkCmdExecuteCommands-pStencilAttachment-06775  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) was
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `stencilAttachmentFormat` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  `VK_FORMAT_UNDEFINED`

- <a href="#VUID-vkCmdExecuteCommands-viewMask-06031"
  id="VUID-vkCmdExecuteCommands-viewMask-06031"></a>
  VUID-vkCmdExecuteCommands-viewMask-06031  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  the `viewMask` member of the
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`viewMask`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdExecuteCommands-pNext-06032"
  id="VUID-vkCmdExecuteCommands-pNext-06032"></a>
  VUID-vkCmdExecuteCommands-pNext-06032  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the `pNext` chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  includes a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, if the `imageView` member of an element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` parameter
  to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the corresponding element of
  the `pColorAttachmentSamples` member of the
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the sample count used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pNext-06033"
  id="VUID-vkCmdExecuteCommands-pNext-06033"></a>
  VUID-vkCmdExecuteCommands-pNext-06033  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the `pNext` chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  includes a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, if the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `depthStencilAttachmentSamples` member of the
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the sample count used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pNext-06034"
  id="VUID-vkCmdExecuteCommands-pNext-06034"></a>
  VUID-vkCmdExecuteCommands-pNext-06034  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the `pNext` chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  includes a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, if the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of the
  `depthStencilAttachmentSamples` member of the
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure included in the `pNext` chain of
  [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`pInheritanceInfo`
  used to begin recording each element of `pCommandBuffers` **must** be
  equal to the sample count used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pNext-06035"
  id="VUID-vkCmdExecuteCommands-pNext-06035"></a>
  VUID-vkCmdExecuteCommands-pNext-06035  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the `pNext` chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  does not include a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, if the `imageView` member of an element of the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments` parameter
  to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`rasterizationSamples`
  **must** be equal to the sample count used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pNext-06036"
  id="VUID-vkCmdExecuteCommands-pNext-06036"></a>
  VUID-vkCmdExecuteCommands-pNext-06036  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the `pNext` chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  does not include a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, if the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`rasterizationSamples`
  **must** be equal to the sample count used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pNext-06037"
  id="VUID-vkCmdExecuteCommands-pNext-06037"></a>
  VUID-vkCmdExecuteCommands-pNext-06037  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  and the `pNext` chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  does not include a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, if the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment->imageView`
  parameter to [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`rasterizationSamples`
  **must** be equal to the sample count used to create that image view

- <a href="#VUID-vkCmdExecuteCommands-pNext-09299"
  id="VUID-vkCmdExecuteCommands-pNext-09299"></a>
  VUID-vkCmdExecuteCommands-pNext-09299  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  with any color attachment using a resolve mode of
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, the `pNext`
  chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  used to create each element of `pCommandBuffers` **must** include a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure with
  an `externalFormat` matching that used to create the resolve
  attachment in the render pass

- <a href="#VUID-vkCmdExecuteCommands-pNext-09300"
  id="VUID-vkCmdExecuteCommands-pNext-09300"></a>
  VUID-vkCmdExecuteCommands-pNext-09300  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  with any color attachment using a resolve mode of
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, and the `pNext`
  chain of
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
  does not include a
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  structure, the value of
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`rasterizationSamples`
  **must** be `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-09375"
  id="VUID-vkCmdExecuteCommands-commandBuffer-09375"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-09375  
  `commandBuffer` **must** not be a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">secondary command buffer</a> unless the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  feature is enabled

- <a href="#VUID-vkCmdExecuteCommands-nestedCommandBuffer-09376"
  id="VUID-vkCmdExecuteCommands-nestedCommandBuffer-09376"></a>
  VUID-vkCmdExecuteCommands-nestedCommandBuffer-09376  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">command buffer nesting level</a> of
  each element of `pCommandBuffers` **must** be less than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxCommandBufferNestingLevel"
  target="_blank"
  rel="noopener"><code>maxCommandBufferNestingLevel</code></a>

- <a href="#VUID-vkCmdExecuteCommands-nestedCommandBufferRendering-09377"
  id="VUID-vkCmdExecuteCommands-nestedCommandBufferRendering-09377"></a>
  VUID-vkCmdExecuteCommands-nestedCommandBufferRendering-09377  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBufferRendering"
  target="_blank"
  rel="noopener"><code>nestedCommandBufferRendering</code></a> feature
  is not enabled, and `commandBuffer` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">secondary command buffer</a>,
  `commandBuffer` **must** not have been recorded with
  `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT`

- <a
  href="#VUID-vkCmdExecuteCommands-nestedCommandBufferSimultaneousUse-09378"
  id="VUID-vkCmdExecuteCommands-nestedCommandBufferSimultaneousUse-09378"></a>
  VUID-vkCmdExecuteCommands-nestedCommandBufferSimultaneousUse-09378  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBufferSimultaneousUse"
  target="_blank"
  rel="noopener"><code>nestedCommandBufferSimultaneousUse</code></a>
  feature is not enabled, and `commandBuffer` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary"
  target="_blank" rel="noopener">secondary command buffer</a>, each
  element of `pCommandBuffers` **must** not have been recorded with
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-09504"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-09504"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-09504  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  the color attachment mapping state specified by
  [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html)
  in the inheritance info of each element of `pCommandBuffers` and in
  the current state of `commandBuffer` **must** match

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-09505"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-09505"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-09505  
  If `vkCmdExecuteCommands` is being called within a render pass
  instance begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
  the input attachment mapping state specified by
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
  in the inheritance info of each element of `pCommandBuffers` and in
  the current state of `commandBuffer` **must** match

Valid Usage (Implicit)

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-parameter"
  id="VUID-vkCmdExecuteCommands-commandBuffer-parameter"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdExecuteCommands-pCommandBuffers-parameter"
  id="VUID-vkCmdExecuteCommands-pCommandBuffers-parameter"></a>
  VUID-vkCmdExecuteCommands-pCommandBuffers-parameter  
  `pCommandBuffers` **must** be a valid pointer to an array of
  `commandBufferCount` valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)
  handles

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-recording"
  id="VUID-vkCmdExecuteCommands-commandBuffer-recording"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdExecuteCommands-commandBuffer-cmdpool"
  id="VUID-vkCmdExecuteCommands-commandBuffer-cmdpool"></a>
  VUID-vkCmdExecuteCommands-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdExecuteCommands-videocoding"
  id="VUID-vkCmdExecuteCommands-videocoding"></a>
  VUID-vkCmdExecuteCommands-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdExecuteCommands-commandBufferCount-arraylength"
  id="VUID-vkCmdExecuteCommands-commandBufferCount-arraylength"></a>
  VUID-vkCmdExecuteCommands-commandBufferCount-arraylength  
  `commandBufferCount` **must** be greater than `0`

- <a href="#VUID-vkCmdExecuteCommands-commonparent"
  id="VUID-vkCmdExecuteCommands-commonparent"></a>
  VUID-vkCmdExecuteCommands-commonparent  
  Both of `commandBuffer`, and the elements of `pCommandBuffers`
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Indirection</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdExecuteCommands"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
