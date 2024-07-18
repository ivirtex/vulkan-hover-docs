# VkRenderPassCreateInfo2(3) Manual Page

## Name

VkRenderPassCreateInfo2 - Structure specifying parameters of a newly
created render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassCreateInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkRenderPassCreateInfo2 {
    VkStructureType                    sType;
    const void*                        pNext;
    VkRenderPassCreateFlags            flags;
    uint32_t                           attachmentCount;
    const VkAttachmentDescription2*    pAttachments;
    uint32_t                           subpassCount;
    const VkSubpassDescription2*       pSubpasses;
    uint32_t                           dependencyCount;
    const VkSubpassDependency2*        pDependencies;
    uint32_t                           correlatedViewMaskCount;
    const uint32_t*                    pCorrelatedViewMasks;
} VkRenderPassCreateInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkRenderPassCreateInfo2 VkRenderPassCreateInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `attachmentCount` is the number of attachments used by this render
  pass.

- `pAttachments` is a pointer to an array of `attachmentCount`
  [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html) structures
  describing the attachments used by the render pass.

- `subpassCount` is the number of subpasses to create.

- `pSubpasses` is a pointer to an array of `subpassCount`
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) structures
  describing each subpass.

- `dependencyCount` is the number of dependencies between pairs of
  subpasses.

- `pDependencies` is a pointer to an array of `dependencyCount`
  [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html) structures
  describing dependencies between pairs of subpasses.

- `correlatedViewMaskCount` is the number of correlation masks.

- `pCorrelatedViewMasks` is a pointer to an array of view masks
  indicating sets of views that **may** be more efficient to render
  concurrently.

## <a href="#_description" class="anchor"></a>Description

Parameters defined by this structure with the same name as those in
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) have the identical
effect to those parameters; the child structures are variants of those
used in [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) which add
`sType` and `pNext` parameters, allowing them to be extended.

If the [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html)::`viewMask`
member of any element of `pSubpasses` is not zero, *multiview*
functionality is considered to be enabled for this render pass.

`correlatedViewMaskCount` and `pCorrelatedViewMasks` have the same
effect as
[VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`correlationMaskCount`
and
[VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pCorrelationMasks`,
respectively.

Valid Usage

- <a href="#VUID-VkRenderPassCreateInfo2-None-03049"
  id="VUID-VkRenderPassCreateInfo2-None-03049"></a>
  VUID-VkRenderPassCreateInfo2-None-03049  
  If any two subpasses operate on attachments with overlapping ranges of
  the same `VkDeviceMemory` object, and at least one subpass writes to
  that area of `VkDeviceMemory`, a subpass dependency **must** be
  included (either directly or via some intermediate subpasses) between
  them

- <a href="#VUID-VkRenderPassCreateInfo2-attachment-03050"
  id="VUID-VkRenderPassCreateInfo2-attachment-03050"></a>
  VUID-VkRenderPassCreateInfo2-attachment-03050  
  If the `attachment` member of any element of `pInputAttachments`,
  `pColorAttachments`, `pResolveAttachments` or
  `pDepthStencilAttachment`, or the attachment indexed by any element of
  `pPreserveAttachments` in any element of `pSubpasses` is bound to a
  range of a `VkDeviceMemory` object that overlaps with any other
  attachment in any subpass (including the same subpass), the
  `VkAttachmentDescription2` structures describing them **must** include
  `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` in `flags`

- <a href="#VUID-VkRenderPassCreateInfo2-attachment-03051"
  id="VUID-VkRenderPassCreateInfo2-attachment-03051"></a>
  VUID-VkRenderPassCreateInfo2-attachment-03051  
  If the `attachment` member of any element of `pInputAttachments`,
  `pColorAttachments`, `pResolveAttachments` or
  `pDepthStencilAttachment`, or any element of `pPreserveAttachments` in
  any element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED`, then it
  **must** be less than `attachmentCount`

- <a
  href="#VUID-VkRenderPassCreateInfo2-fragmentDensityMapAttachment-06472"
  id="VUID-VkRenderPassCreateInfo2-fragmentDensityMapAttachment-06472"></a>
  VUID-VkRenderPassCreateInfo2-fragmentDensityMapAttachment-06472  
  If the pNext chain includes a
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)
  structure and the `fragmentDensityMapAttachment` member is not
  `VK_ATTACHMENT_UNUSED`, then `attachment` **must** be less than
  `attachmentCount`

- <a href="#VUID-VkRenderPassCreateInfo2-pSubpasses-06473"
  id="VUID-VkRenderPassCreateInfo2-pSubpasses-06473"></a>
  VUID-VkRenderPassCreateInfo2-pSubpasses-06473  
  If the `pSubpasses` pNext chain includes a
  [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html)
  structure and the `pDepthStencilResolveAttachment` member is not
  `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, then
  `attachment` **must** be less than `attachmentCount`

- <a href="#VUID-VkRenderPassCreateInfo2-pAttachments-02522"
  id="VUID-VkRenderPassCreateInfo2-pAttachments-02522"></a>
  VUID-VkRenderPassCreateInfo2-pAttachments-02522  
  For any member of `pAttachments` with a `loadOp` equal to
  `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment
  **must** not specify a `layout` equal to
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderPassCreateInfo2-pAttachments-02523"
  id="VUID-VkRenderPassCreateInfo2-pAttachments-02523"></a>
  VUID-VkRenderPassCreateInfo2-pAttachments-02523  
  For any member of `pAttachments` with a `stencilLoadOp` equal to
  `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment
  **must** not specify a `layout` equal to
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderPassCreateInfo2-pDependencies-03054"
  id="VUID-VkRenderPassCreateInfo2-pDependencies-03054"></a>
  VUID-VkRenderPassCreateInfo2-pDependencies-03054  
  For any element of `pDependencies`, if the `srcSubpass` is not
  `VK_SUBPASS_EXTERNAL`, all stage flags included in the `srcStageMask`
  member of that dependency **must** be a pipeline stage supported by
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-types"
  target="_blank" rel="noopener">pipeline</a> identified by the
  `pipelineBindPoint` member of the source subpass

- <a href="#VUID-VkRenderPassCreateInfo2-pDependencies-03055"
  id="VUID-VkRenderPassCreateInfo2-pDependencies-03055"></a>
  VUID-VkRenderPassCreateInfo2-pDependencies-03055  
  For any element of `pDependencies`, if the `dstSubpass` is not
  `VK_SUBPASS_EXTERNAL`, all stage flags included in the `dstStageMask`
  member of that dependency **must** be a pipeline stage supported by
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-types"
  target="_blank" rel="noopener">pipeline</a> identified by the
  `pipelineBindPoint` member of the destination subpass

- <a href="#VUID-VkRenderPassCreateInfo2-pCorrelatedViewMasks-03056"
  id="VUID-VkRenderPassCreateInfo2-pCorrelatedViewMasks-03056"></a>
  VUID-VkRenderPassCreateInfo2-pCorrelatedViewMasks-03056  
  The set of bits included in any element of `pCorrelatedViewMasks`
  **must** not overlap with the set of bits included in any other
  element of `pCorrelatedViewMasks`

- <a href="#VUID-VkRenderPassCreateInfo2-viewMask-03057"
  id="VUID-VkRenderPassCreateInfo2-viewMask-03057"></a>
  VUID-VkRenderPassCreateInfo2-viewMask-03057  
  If the [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html)::`viewMask`
  member of all elements of `pSubpasses` is `0`,
  `correlatedViewMaskCount` **must** be `0`

- <a href="#VUID-VkRenderPassCreateInfo2-viewMask-03058"
  id="VUID-VkRenderPassCreateInfo2-viewMask-03058"></a>
  VUID-VkRenderPassCreateInfo2-viewMask-03058  
  The [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html)::`viewMask`
  member of all elements of `pSubpasses` **must** either all be `0`, or
  all not be `0`

- <a href="#VUID-VkRenderPassCreateInfo2-viewMask-03059"
  id="VUID-VkRenderPassCreateInfo2-viewMask-03059"></a>
  VUID-VkRenderPassCreateInfo2-viewMask-03059  
  If the [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html)::`viewMask`
  member of all elements of `pSubpasses` is `0`, the `dependencyFlags`
  member of any element of `pDependencies` **must** not include
  `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-VkRenderPassCreateInfo2-pDependencies-03060"
  id="VUID-VkRenderPassCreateInfo2-pDependencies-03060"></a>
  VUID-VkRenderPassCreateInfo2-pDependencies-03060  
  For any element of `pDependencies` where its `srcSubpass` member
  equals its `dstSubpass` member, if the `viewMask` member of the
  corresponding element of `pSubpasses` includes more than one bit, its
  `dependencyFlags` member **must** include
  `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-VkRenderPassCreateInfo2-attachment-02525"
  id="VUID-VkRenderPassCreateInfo2-attachment-02525"></a>
  VUID-VkRenderPassCreateInfo2-attachment-02525  
  If the `attachment` member of any element of the `pInputAttachments`
  member of any element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED`,
  the `aspectMask` member of that element of `pInputAttachments`
  **must** only include aspects that are present in images of the format
  specified by the element of `pAttachments` specified by `attachment`

- <a href="#VUID-VkRenderPassCreateInfo2-srcSubpass-02526"
  id="VUID-VkRenderPassCreateInfo2-srcSubpass-02526"></a>
  VUID-VkRenderPassCreateInfo2-srcSubpass-02526  
  The `srcSubpass` member of each element of `pDependencies` **must** be
  less than `subpassCount`

- <a href="#VUID-VkRenderPassCreateInfo2-dstSubpass-02527"
  id="VUID-VkRenderPassCreateInfo2-dstSubpass-02527"></a>
  VUID-VkRenderPassCreateInfo2-dstSubpass-02527  
  The `dstSubpass` member of each element of `pDependencies` **must** be
  less than `subpassCount`

- <a href="#VUID-VkRenderPassCreateInfo2-pAttachments-04585"
  id="VUID-VkRenderPassCreateInfo2-pAttachments-04585"></a>
  VUID-VkRenderPassCreateInfo2-pAttachments-04585  
  If any element of `pAttachments` is used as a fragment shading rate
  attachment in any subpass, it **must** not be used as any other
  attachment in the render pass

- <a href="#VUID-VkRenderPassCreateInfo2-pAttachments-09387"
  id="VUID-VkRenderPassCreateInfo2-pAttachments-09387"></a>
  VUID-VkRenderPassCreateInfo2-pAttachments-09387  
  If any element of `pAttachments` is used as a fragment shading rate
  attachment, the `loadOp` for that attachment **must** not be
  `VK_ATTACHMENT_LOAD_OP_CLEAR`

- <a href="#VUID-VkRenderPassCreateInfo2-flags-04521"
  id="VUID-VkRenderPassCreateInfo2-flags-04521"></a>
  VUID-VkRenderPassCreateInfo2-flags-04521  
  If `flags` includes `VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM`, an
  element of `pSubpasses` includes an instance of
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)
  in its `pNext` chain, and the `pFragmentShadingRateAttachment` member
  of that structure is not equal to `NULL`, the `attachment` member of
  `pFragmentShadingRateAttachment` **must** be `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkRenderPassCreateInfo2-pAttachments-04586"
  id="VUID-VkRenderPassCreateInfo2-pAttachments-04586"></a>
  VUID-VkRenderPassCreateInfo2-pAttachments-04586  
  If any element of `pAttachments` is used as a fragment shading rate
  attachment in any subpass, it **must** have an image format whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> contain
  `VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkRenderPassCreateInfo2-rasterizationSamples-04905"
  id="VUID-VkRenderPassCreateInfo2-rasterizationSamples-04905"></a>
  VUID-VkRenderPassCreateInfo2-rasterizationSamples-04905  
  If the pipeline is being created with fragment shader state, and the
  [`VK_QCOM_render_pass_shader_resolve`](VK_QCOM_render_pass_shader_resolve.html)` extension`
  is enabled, and if subpass has any input attachments, and if the
  subpass description contains
  `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`, then the sample
  count of the input attachments **must** equal `rasterizationSamples`

- <a href="#VUID-VkRenderPassCreateInfo2-sampleShadingEnable-04906"
  id="VUID-VkRenderPassCreateInfo2-sampleShadingEnable-04906"></a>
  VUID-VkRenderPassCreateInfo2-sampleShadingEnable-04906  
  If the pipeline is being created with fragment shader state, and the
  [`VK_QCOM_render_pass_shader_resolve`](VK_QCOM_render_pass_shader_resolve.html)
  extension is enabled, and if the subpass description contains
  `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`, then
  `sampleShadingEnable` **must** be false

- <a href="#VUID-VkRenderPassCreateInfo2-flags-04907"
  id="VUID-VkRenderPassCreateInfo2-flags-04907"></a>
  VUID-VkRenderPassCreateInfo2-flags-04907  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`,
  and if `pResolveAttachments` is not `NULL`, then each resolve
  attachment **must** be `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkRenderPassCreateInfo2-flags-04908"
  id="VUID-VkRenderPassCreateInfo2-flags-04908"></a>
  VUID-VkRenderPassCreateInfo2-flags-04908  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`,
  and if `pDepthStencilResolveAttachment` is not `NULL`, then the
  depth/stencil resolve attachment **must** be `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkRenderPassCreateInfo2-flags-04909"
  id="VUID-VkRenderPassCreateInfo2-flags-04909"></a>
  VUID-VkRenderPassCreateInfo2-flags-04909  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`,
  then the subpass **must** be the last subpass in a subpass dependency
  chain

- <a href="#VUID-VkRenderPassCreateInfo2-attachment-06244"
  id="VUID-VkRenderPassCreateInfo2-attachment-06244"></a>
  VUID-VkRenderPassCreateInfo2-attachment-06244  
  If the `attachment` member of the `pDepthStencilAttachment` member of
  an element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED`, the `layout`
  member of that same structure is either
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, and the `pNext` chain of
  that structure does not include a
  [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayout.html)
  structure, then the element of `pAttachments` with an index equal to
  `attachment` **must** not have a `format` that includes both depth and
  stencil components

- <a href="#VUID-VkRenderPassCreateInfo2-attachment-06245"
  id="VUID-VkRenderPassCreateInfo2-attachment-06245"></a>
  VUID-VkRenderPassCreateInfo2-attachment-06245  
  If the `attachment` member of the `pDepthStencilAttachment` member of
  an element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED` and the
  `layout` member of that same structure is either
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`, then the element of
  `pAttachments` with an index equal to `attachment` **must** have a
  `format` that includes only a stencil component

- <a href="#VUID-VkRenderPassCreateInfo2-attachment-06246"
  id="VUID-VkRenderPassCreateInfo2-attachment-06246"></a>
  VUID-VkRenderPassCreateInfo2-attachment-06246  
  If the `attachment` member of the `pDepthStencilAttachment` member of
  an element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED` and the
  `layout` member of that same structure is either
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, then the element of
  `pAttachments` with an index equal to `attachment` **must** not have a
  `format` that includes only a stencil component

- <a href="#VUID-VkRenderPassCreateInfo2-pResolveAttachments-09331"
  id="VUID-VkRenderPassCreateInfo2-pResolveAttachments-09331"></a>
  VUID-VkRenderPassCreateInfo2-pResolveAttachments-09331  
  If any element of `pResolveAttachments` of any element of `pSubpasses`
  references an attachment description with a format of
  `VK_FORMAT_UNDEFINED`,
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment->attachment`
  **must** be `VK_ATTACHMENT_UNUSED`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassCreateInfo2-sType-sType"
  id="VUID-VkRenderPassCreateInfo2-sType-sType"></a>
  VUID-VkRenderPassCreateInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2`

- <a href="#VUID-VkRenderPassCreateInfo2-pNext-pNext"
  id="VUID-VkRenderPassCreateInfo2-pNext-pNext"></a>
  VUID-VkRenderPassCreateInfo2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationControlEXT.html),
  [VkRenderPassCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackCreateInfoEXT.html),
  or
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)

- <a href="#VUID-VkRenderPassCreateInfo2-sType-unique"
  id="VUID-VkRenderPassCreateInfo2-sType-unique"></a>
  VUID-VkRenderPassCreateInfo2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkRenderPassCreateInfo2-flags-parameter"
  id="VUID-VkRenderPassCreateInfo2-flags-parameter"></a>
  VUID-VkRenderPassCreateInfo2-flags-parameter  
  `flags` **must** be a valid combination of
  [VkRenderPassCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlagBits.html) values

- <a href="#VUID-VkRenderPassCreateInfo2-pAttachments-parameter"
  id="VUID-VkRenderPassCreateInfo2-pAttachments-parameter"></a>
  VUID-VkRenderPassCreateInfo2-pAttachments-parameter  
  If `attachmentCount` is not `0`, `pAttachments` **must** be a valid
  pointer to an array of `attachmentCount` valid
  [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html) structures

- <a href="#VUID-VkRenderPassCreateInfo2-pSubpasses-parameter"
  id="VUID-VkRenderPassCreateInfo2-pSubpasses-parameter"></a>
  VUID-VkRenderPassCreateInfo2-pSubpasses-parameter  
  `pSubpasses` **must** be a valid pointer to an array of `subpassCount`
  valid [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) structures

- <a href="#VUID-VkRenderPassCreateInfo2-pDependencies-parameter"
  id="VUID-VkRenderPassCreateInfo2-pDependencies-parameter"></a>
  VUID-VkRenderPassCreateInfo2-pDependencies-parameter  
  If `dependencyCount` is not `0`, `pDependencies` **must** be a valid
  pointer to an array of `dependencyCount` valid
  [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html) structures

- <a href="#VUID-VkRenderPassCreateInfo2-pCorrelatedViewMasks-parameter"
  id="VUID-VkRenderPassCreateInfo2-pCorrelatedViewMasks-parameter"></a>
  VUID-VkRenderPassCreateInfo2-pCorrelatedViewMasks-parameter  
  If `correlatedViewMaskCount` is not `0`, `pCorrelatedViewMasks`
  **must** be a valid pointer to an array of `correlatedViewMaskCount`
  `uint32_t` values

- <a href="#VUID-VkRenderPassCreateInfo2-subpassCount-arraylength"
  id="VUID-VkRenderPassCreateInfo2-subpassCount-arraylength"></a>
  VUID-VkRenderPassCreateInfo2-subpassCount-arraylength  
  `subpassCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html),
[VkRenderPassCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html),
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html),
[vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html),
[vkCreateRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassCreateInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
