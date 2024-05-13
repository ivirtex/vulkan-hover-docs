# VkImageViewCreateInfo(3) Manual Page

## Name

VkImageViewCreateInfo - Structure specifying parameters of a newly
created image view



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageViewCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageViewCreateInfo {
    VkStructureType            sType;
    const void*                pNext;
    VkImageViewCreateFlags     flags;
    VkImage                    image;
    VkImageViewType            viewType;
    VkFormat                   format;
    VkComponentMapping         components;
    VkImageSubresourceRange    subresourceRange;
} VkImageViewCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlagBits.html) specifying
  additional parameters of the image view.

- `image` is a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) on which the view will be
  created.

- `viewType` is a [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) value
  specifying the type of the image view.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) specifying the format and type
  used to interpret texel blocks of the image.

- `components` is a [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html)
  structure specifying a remapping of color components (or of depth or
  stencil components after they have been converted into color
  components).

- `subresourceRange` is a
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html) structure
  selecting the set of mipmap levels and array layers to be accessible
  to the view.

## <a href="#_description" class="anchor"></a>Description

Some of the `image` creation parameters are inherited by the view. In
particular, image view creation inherits the implicit parameter `usage`
specifying the allowed usages of the image view that, by default, takes
the value of the corresponding `usage` parameter specified in
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) at image creation time. The
implicit `usage` **can** be overridden by adding a
[VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html) structure
to the `pNext` chain, but the view usage **must** be a subset of the
image usage. If `image` has a depth-stencil format and was created with
a [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
structure included in the `pNext` chain of
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), the usage is calculated
based on the `subresource.aspectMask` provided:

- If `aspectMask` includes only `VK_IMAGE_ASPECT_STENCIL_BIT`, the
  implicit `usage` is equal to
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`.

- If `aspectMask` includes only `VK_IMAGE_ASPECT_DEPTH_BIT`, the
  implicit `usage` is equal to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`.

- If both aspects are included in `aspectMask`, the implicit `usage` is
  equal to the intersection of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` and
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`.

If `image` is a 3D image, its Z range **can** be restricted to a subset
by adding a
[VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSlicedCreateInfoEXT.html) to
the `pNext` chain.

If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`
flag, and if the `format` of the image is not <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar</a>, `format` **can** be
different from the image’s format, but if `image` was created without
the `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag and they are
not equal they **must** be *compatible*. Image format compatibility is
defined in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility-classes"
target="_blank" rel="noopener">Format Compatibility Classes</a> section.
Views of compatible formats will have the same mapping between texel
coordinates and memory locations irrespective of the `format`, with only
the interpretation of the bit pattern changing.

If `image` was created with a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar</a> format, and the image
view’s `aspectMask` is one of `VK_IMAGE_ASPECT_PLANE_0_BIT`,
`VK_IMAGE_ASPECT_PLANE_1_BIT` or `VK_IMAGE_ASPECT_PLANE_2_BIT`, the
view’s aspect mask is considered to be equivalent to
`VK_IMAGE_ASPECT_COLOR_BIT` when used as a framebuffer attachment.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Values intended to be used with one view format <strong>may</strong>
not be exactly preserved when written or read through a different
format. For example, an integer value that happens to have the bit
pattern of a floating point denorm or NaN <strong>may</strong> be
flushed or canonicalized when written or read through a view with a
floating point format. Similarly, a value written through a signed
normalized format that has a bit pattern exactly equal to -2<sup>b</sup>
<strong>may</strong> be changed to -2<sup>b</sup> + 1 as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-fixedfpconv"
target="_blank" rel="noopener">Conversion from Normalized Fixed-Point to
Floating-Point</a>.</p></td>
</tr>
</tbody>
</table>

If `image` was created with the
`VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, `format`
**must** be *compatible* with the image’s format as described above; or
**must** be an uncompressed format, in which case it **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-size-compatibility"
target="_blank" rel="noopener"><em>size-compatible</em></a> with the
image’s format. In this case, the resulting image view’s texel
dimensions equal the dimensions of the selected mip level divided by the
compressed texel block size and rounded up.

The [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html) `components` member
describes a remapping from components of the image to components of the
vector returned by shader image instructions. This remapping **must** be
the identity swizzle for storage image descriptors, input attachment
descriptors, framebuffer attachments, and any `VkImageView` used with a
combined image sampler that enables <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
conversion</a>.

If the image view is to be used with a sampler which supports <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
conversion</a>, an *identically defined object* of type
[VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) to that used
to create the sampler **must** be passed to
[vkCreateImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html) in a
[VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
included in the `pNext` chain of
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html). Conversely, if a
[VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) object is
passed to [vkCreateImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html), an identically
defined [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) object
**must** be used when sampling the image.

If the image has a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar</a> `format`,
`subresourceRange.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, and
`usage` includes `VK_IMAGE_USAGE_SAMPLED_BIT`, then the `format`
**must** be identical to the image `format` and the sampler to be used
with the image view **must** enable <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
conversion</a>.

When such an image is used in a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding"
target="_blank" rel="noopener">video coding</a> operation, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
conversion</a> has no effect.

If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` and
the image has a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">multi-planar</a> `format`, and if
`subresourceRange.aspectMask` is `VK_IMAGE_ASPECT_PLANE_0_BIT`,
`VK_IMAGE_ASPECT_PLANE_1_BIT`, or `VK_IMAGE_ASPECT_PLANE_2_BIT`,
`format` **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatible-planes"
target="_blank" rel="noopener">compatible</a> with the corresponding
plane of the image, and the sampler to be used with the image view
**must** not enable <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
conversion</a>. The `width` and `height` of the single-plane image view
**must** be derived from the multi-planar image’s dimensions in the
manner listed for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatible-planes"
target="_blank" rel="noopener">plane compatibility</a> for the plane.

Any view of an image plane will have the same mapping between texel
coordinates and memory locations as used by the components of the color
aspect, subject to the formulae relating texel coordinates to
lower-resolution planes as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-chroma-reconstruction"
target="_blank" rel="noopener">Chroma Reconstruction</a>. That is, if an
R or B plane has a reduced resolution relative to the G plane of the
multi-planar image, the image view operates using the
(*u<sub>plane</sub>*, *v<sub>plane</sub>*) unnormalized coordinates of
the reduced-resolution plane, and these coordinates access the same
memory locations as the (*u<sub>color</sub>*, *v<sub>color</sub>*)
unnormalized coordinates of the color aspect for which chroma
reconstruction operations operate on the same (*u<sub>plane</sub>*,
*v<sub>plane</sub>*) or (*i<sub>plane</sub>*, *j<sub>plane</sub>*)
coordinates.

| Image View Type                 | Compatible Image Types                  |
|---------------------------------|-----------------------------------------|
| `VK_IMAGE_VIEW_TYPE_1D`         | `VK_IMAGE_TYPE_1D`                      |
| `VK_IMAGE_VIEW_TYPE_1D_ARRAY`   | `VK_IMAGE_TYPE_1D`                      |
| `VK_IMAGE_VIEW_TYPE_2D`         | `VK_IMAGE_TYPE_2D` , `VK_IMAGE_TYPE_3D` |
| `VK_IMAGE_VIEW_TYPE_2D_ARRAY`   | `VK_IMAGE_TYPE_2D` , `VK_IMAGE_TYPE_3D` |
| `VK_IMAGE_VIEW_TYPE_CUBE`       | `VK_IMAGE_TYPE_2D`                      |
| `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY` | `VK_IMAGE_TYPE_2D`                      |
| `VK_IMAGE_VIEW_TYPE_3D`         | `VK_IMAGE_TYPE_3D`                      |

Table 1. Image type and image view type compatibility requirements
{#resources-image-views-compatibility}

Valid Usage

- <a href="#VUID-VkImageViewCreateInfo-image-01003"
  id="VUID-VkImageViewCreateInfo-image-01003"></a>
  VUID-VkImageViewCreateInfo-image-01003  
  If `image` was not created with `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`
  then `viewType` **must** not be `VK_IMAGE_VIEW_TYPE_CUBE` or
  `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-viewType-01004"
  id="VUID-VkImageViewCreateInfo-viewType-01004"></a>
  VUID-VkImageViewCreateInfo-viewType-01004  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageCubeArray"
  target="_blank" rel="noopener"><code>imageCubeArray</code></a> feature
  is not enabled, `viewType` **must** not be
  `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-image-06723"
  id="VUID-VkImageViewCreateInfo-image-06723"></a>
  VUID-VkImageViewCreateInfo-image-06723  
  If `image` was created with `VK_IMAGE_TYPE_3D` but without
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set then `viewType` **must**
  not be `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-image-06728"
  id="VUID-VkImageViewCreateInfo-image-06728"></a>
  VUID-VkImageViewCreateInfo-image-06728  
  If `image` was created with `VK_IMAGE_TYPE_3D` but without
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` or
  `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set, then `viewType`
  **must** not be `VK_IMAGE_VIEW_TYPE_2D`

- <a href="#VUID-VkImageViewCreateInfo-image-04970"
  id="VUID-VkImageViewCreateInfo-image-04970"></a>
  VUID-VkImageViewCreateInfo-image-04970  
  If `image` was created with `VK_IMAGE_TYPE_3D` and `viewType` is
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY` then
  `subresourceRange.levelCount` **must** be 1

- <a href="#VUID-VkImageViewCreateInfo-image-04971"
  id="VUID-VkImageViewCreateInfo-image-04971"></a>
  VUID-VkImageViewCreateInfo-image-04971  
  If `image` was created with `VK_IMAGE_TYPE_3D` and `viewType` is
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY` then
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` **must** not
  contain any of `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`,
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, and
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`

- <a href="#VUID-VkImageViewCreateInfo-image-04972"
  id="VUID-VkImageViewCreateInfo-image-04972"></a>
  VUID-VkImageViewCreateInfo-image-04972  
  If `image` was created with a `samples` value not equal to
  `VK_SAMPLE_COUNT_1_BIT` then `viewType` **must** be either
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-image-04441"
  id="VUID-VkImageViewCreateInfo-image-04441"></a>
  VUID-VkImageViewCreateInfo-image-04441  
  `image` **must** have been created with a `usage` value containing at
  least one of the usages defined in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#valid-imageview-imageusage"
  target="_blank" rel="noopener">valid image usage</a> list for image
  views

- <a href="#VUID-VkImageViewCreateInfo-None-02273"
  id="VUID-VkImageViewCreateInfo-None-02273"></a>
  VUID-VkImageViewCreateInfo-None-02273  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> of the resultant
  image view **must** contain at least one bit

- <a href="#VUID-VkImageViewCreateInfo-usage-02274"
  id="VUID-VkImageViewCreateInfo-usage-02274"></a>
  VUID-VkImageViewCreateInfo-usage-02274  
  If `usage` contains `VK_IMAGE_USAGE_SAMPLED_BIT`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> of the resultant
  image view **must** contain `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`

- <a href="#VUID-VkImageViewCreateInfo-usage-02275"
  id="VUID-VkImageViewCreateInfo-usage-02275"></a>
  VUID-VkImageViewCreateInfo-usage-02275  
  If `usage` contains `VK_IMAGE_USAGE_STORAGE_BIT`, then the image
  view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT`

- <a href="#VUID-VkImageViewCreateInfo-usage-08931"
  id="VUID-VkImageViewCreateInfo-usage-08931"></a>
  VUID-VkImageViewCreateInfo-usage-08931  
  If `usage` contains `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`, then the
  image view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or
  `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkImageViewCreateInfo-usage-02277"
  id="VUID-VkImageViewCreateInfo-usage-02277"></a>
  VUID-VkImageViewCreateInfo-usage-02277  
  If `usage` contains `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  then the image view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageViewCreateInfo-image-08333"
  id="VUID-VkImageViewCreateInfo-image-08333"></a>
  VUID-VkImageViewCreateInfo-image-08333  
  If `image` was created with
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` and `usage`
  contains `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, then the image
  view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_VIDEO_DECODE_OUTPUT_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-image-08334"
  id="VUID-VkImageViewCreateInfo-image-08334"></a>
  VUID-VkImageViewCreateInfo-image-08334  
  If `image` was created with
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` and `usage`
  contains `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, then the image
  view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_VIDEO_DECODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-image-08335"
  id="VUID-VkImageViewCreateInfo-image-08335"></a>
  VUID-VkImageViewCreateInfo-image-08335  
  If `image` was created with
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then `usage`
  **must** not include `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-image-08336"
  id="VUID-VkImageViewCreateInfo-image-08336"></a>
  VUID-VkImageViewCreateInfo-image-08336  
  If `image` was created with
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` and `usage`
  contains `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`, then the image
  view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_VIDEO_ENCODE_INPUT_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-image-08337"
  id="VUID-VkImageViewCreateInfo-image-08337"></a>
  VUID-VkImageViewCreateInfo-image-08337  
  If `image` was created with
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` and `usage`
  contains `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, then the image
  view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_VIDEO_ENCODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-image-08338"
  id="VUID-VkImageViewCreateInfo-image-08338"></a>
  VUID-VkImageViewCreateInfo-image-08338  
  If `image` was created with
  `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then `usage`
  **must** not include `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-usage-08932"
  id="VUID-VkImageViewCreateInfo-usage-08932"></a>
  VUID-VkImageViewCreateInfo-usage-08932  
  If `usage` contains `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`, and any of
  the following is true:

  - the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
    target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
    feature is not enabled

  - the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
    target="_blank"
    rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
    property is `VK_FALSE`

  - `image` was created with an
    [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
    value of 0

  then the image view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain at
  least one of `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or
  `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT` or
  `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-01478"
  id="VUID-VkImageViewCreateInfo-subresourceRange-01478"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-01478  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels`
  specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image`
  was created

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-01718"
  id="VUID-VkImageViewCreateInfo-subresourceRange-01718"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-01718  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`,
  `subresourceRange.baseMipLevel` + `subresourceRange.levelCount`
  **must** be less than or equal to the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkImageViewCreateInfo-image-02571"
  id="VUID-VkImageViewCreateInfo-image-02571"></a>
  VUID-VkImageViewCreateInfo-image-02571  
  If `image` was created with `usage` containing
  `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`,
  `subresourceRange.levelCount` **must** be `1`

- <a href="#VUID-VkImageViewCreateInfo-image-06724"
  id="VUID-VkImageViewCreateInfo-image-06724"></a>
  VUID-VkImageViewCreateInfo-image-06724  
  If `image` is not a 3D image created with
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` or
  `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set, or `viewType` is not
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`,
  `subresourceRange.baseArrayLayer` **must** be less than the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `image` was created

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-06725"
  id="VUID-VkImageViewCreateInfo-subresourceRange-06725"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-06725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `image` is not a 3D image created with
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` or
  `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT` set, or `viewType` is not
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`,
  `subresourceRange.layerCount` **must** be non-zero and
  `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount`
  **must** be less than or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkImageViewCreateInfo-image-02724"
  id="VUID-VkImageViewCreateInfo-image-02724"></a>
  VUID-VkImageViewCreateInfo-image-02724  
  If `image` is a 3D image created with
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, and `viewType` is
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`,
  `subresourceRange.baseArrayLayer` **must** be less than the depth
  computed from `baseMipLevel` and `extent.depth` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created,
  according to the formula defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
  target="_blank" rel="noopener">Image Mip Level Sizing</a>

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-02725"
  id="VUID-VkImageViewCreateInfo-subresourceRange-02725"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-02725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `image` is a 3D image created with
  `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, and `viewType` is
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`,
  `subresourceRange.layerCount` **must** be non-zero and
  `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount`
  **must** be less than or equal to the depth computed from
  `baseMipLevel` and `extent.depth` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created,
  according to the formula defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
  target="_blank" rel="noopener">Image Mip Level Sizing</a>

- <a href="#VUID-VkImageViewCreateInfo-image-01761"
  id="VUID-VkImageViewCreateInfo-image-01761"></a>
  VUID-VkImageViewCreateInfo-image-01761  
  If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`
  flag, but without the
  `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, and if the
  `format` of the `image` is not a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar</a> format, `format`
  **must** be compatible with the `format` used to create `image`, as
  defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility-classes"
  target="_blank" rel="noopener">Format Compatibility Classes</a>

- <a href="#VUID-VkImageViewCreateInfo-image-01583"
  id="VUID-VkImageViewCreateInfo-image-01583"></a>
  VUID-VkImageViewCreateInfo-image-01583  
  If `image` was created with the
  `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, `format`
  **must** be compatible with, or **must** be an uncompressed format
  that is size-compatible with, the `format` used to create `image`

- <a href="#VUID-VkImageViewCreateInfo-image-07072"
  id="VUID-VkImageViewCreateInfo-image-07072"></a>
  VUID-VkImageViewCreateInfo-image-07072  
  If `image` was created with the
  `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag and `format` is
  a non-compressed format, the `levelCount` member of `subresourceRange`
  **must** be `1`

- <a href="#VUID-VkImageViewCreateInfo-image-09487"
  id="VUID-VkImageViewCreateInfo-image-09487"></a>
  VUID-VkImageViewCreateInfo-image-09487  
  If `image` was created with the
  `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` flag, the
  `VkPhysicalDeviceMaintenance6PropertiesKHR`::`blockTexelViewCompatibleMultipleLayers`
  property is not set to `VK_TRUE`, and `format` is a non-compressed
  format, then the `layerCount` member of `subresourceRange` **must** be
  `1`

- <a href="#VUID-VkImageViewCreateInfo-pNext-01585"
  id="VUID-VkImageViewCreateInfo-pNext-01585"></a>
  VUID-VkImageViewCreateInfo-pNext-01585  
  If a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure was included in the `pNext` chain of the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure used when
  creating `image` and
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  is not zero then `format` **must** be one of the formats in
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats`

- <a href="#VUID-VkImageViewCreateInfo-image-01586"
  id="VUID-VkImageViewCreateInfo-image-01586"></a>
  VUID-VkImageViewCreateInfo-image-01586  
  If `image` was created with the `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`
  flag, if the `format` of the `image` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar</a> format, and if
  `subresourceRange.aspectMask` is one of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-planes-image-aspect"
  target="_blank" rel="noopener">multi-planar aspect mask</a> bits, then
  `format` **must** be compatible with the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) for
  the plane of the `image` `format` indicated by
  `subresourceRange.aspectMask`, as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatible-planes"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatible-planes</a>

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-07818"
  id="VUID-VkImageViewCreateInfo-subresourceRange-07818"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-07818  
  `subresourceRange.aspectMask` **must** only have at most 1 valid <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-planes-image-aspect"
  target="_blank" rel="noopener">multi-planar aspect mask</a> bit

- <a href="#VUID-VkImageViewCreateInfo-image-01762"
  id="VUID-VkImageViewCreateInfo-image-01762"></a>
  VUID-VkImageViewCreateInfo-image-01762  
  If `image` was not created with the
  `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` flag, or if the `format` of the
  `image` is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">multi-planar</a> format and if
  `subresourceRange.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, `format`
  **must** be identical to the `format` used to create `image`

- <a href="#VUID-VkImageViewCreateInfo-format-06415"
  id="VUID-VkImageViewCreateInfo-format-06415"></a>
  VUID-VkImageViewCreateInfo-format-06415  
  If the image view <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#image-views-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">requires a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a> and `usage` contains
  `VK_IMAGE_USAGE_SAMPLED_BIT`, then the `pNext` chain **must** include
  a [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
  structure with a conversion value other than
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkImageViewCreateInfo-format-04714"
  id="VUID-VkImageViewCreateInfo-format-04714"></a>
  VUID-VkImageViewCreateInfo-format-04714  
  If `format` has a `_422` or `_420` suffix then `image` **must** have
  been created with a width that is a multiple of 2

- <a href="#VUID-VkImageViewCreateInfo-format-04715"
  id="VUID-VkImageViewCreateInfo-format-04715"></a>
  VUID-VkImageViewCreateInfo-format-04715  
  If `format` has a `_420` suffix then `image` **must** have been
  created with a height that is a multiple of 2

- <a href="#VUID-VkImageViewCreateInfo-pNext-01970"
  id="VUID-VkImageViewCreateInfo-pNext-01970"></a>
  VUID-VkImageViewCreateInfo-pNext-01970  
  If the `pNext` chain includes a
  [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
  structure with a `conversion` value other than
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), all members of `components`
  **must** have the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkImageViewCreateInfo-pNext-06658"
  id="VUID-VkImageViewCreateInfo-pNext-06658"></a>
  VUID-VkImageViewCreateInfo-pNext-06658  
  If the `pNext` chain includes a
  [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
  structure with a `conversion` value other than
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `format` **must** be the same
  used in
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)::`format`

- <a href="#VUID-VkImageViewCreateInfo-image-01020"
  id="VUID-VkImageViewCreateInfo-image-01020"></a>
  VUID-VkImageViewCreateInfo-image-01020  
  If `image` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkImageViewCreateInfo-subResourceRange-01021"
  id="VUID-VkImageViewCreateInfo-subResourceRange-01021"></a>
  VUID-VkImageViewCreateInfo-subResourceRange-01021  
  `viewType` **must** be compatible with the type of `image` as shown in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-compatibility"
  target="_blank" rel="noopener">view type compatibility table</a>

- <a href="#VUID-VkImageViewCreateInfo-image-02399"
  id="VUID-VkImageViewCreateInfo-image-02399"></a>
  VUID-VkImageViewCreateInfo-image-02399  
  If `image` has an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
  target="_blank" rel="noopener">Android external format</a>, `format`
  **must** be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkImageViewCreateInfo-image-02400"
  id="VUID-VkImageViewCreateInfo-image-02400"></a>
  VUID-VkImageViewCreateInfo-image-02400  
  If `image` has an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
  target="_blank" rel="noopener">Android external format</a>, the
  `pNext` chain **must** include a
  [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
  structure with a `conversion` object created with the same external
  format as `image`

- <a href="#VUID-VkImageViewCreateInfo-image-02401"
  id="VUID-VkImageViewCreateInfo-image-02401"></a>
  VUID-VkImageViewCreateInfo-image-02401  
  If `image` has an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
  target="_blank" rel="noopener">Android external format</a>, all
  members of `components` **must** be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkImageViewCreateInfo-image-08957"
  id="VUID-VkImageViewCreateInfo-image-08957"></a>
  VUID-VkImageViewCreateInfo-image-08957  
  If `image` has an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-external-formats"
  target="_blank" rel="noopener">QNX Screen external format</a>,
  `format` **must** be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkImageViewCreateInfo-image-08958"
  id="VUID-VkImageViewCreateInfo-image-08958"></a>
  VUID-VkImageViewCreateInfo-image-08958  
  If `image` has an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-external-formats"
  target="_blank" rel="noopener">QNX Screen external format</a>, the
  `pNext` chain **must** include a
  [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)
  structure with a `conversion` object created with the same external
  format as `image`

- <a href="#VUID-VkImageViewCreateInfo-image-08959"
  id="VUID-VkImageViewCreateInfo-image-08959"></a>
  VUID-VkImageViewCreateInfo-image-08959  
  If `image` has an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-external-formats"
  target="_blank" rel="noopener">QNX Screen external format</a>, all
  members of `components` **must** be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkImageViewCreateInfo-image-02086"
  id="VUID-VkImageViewCreateInfo-image-02086"></a>
  VUID-VkImageViewCreateInfo-image-02086  
  If `image` was created with `usage` containing
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, `viewType`
  **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-image-02087"
  id="VUID-VkImageViewCreateInfo-image-02087"></a>
  VUID-VkImageViewCreateInfo-image-02087  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature is enabled, and If `image` was created with `usage` containing
  `VK_IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV`, `format` **must** be
  `VK_FORMAT_R8_UINT`

- <a href="#VUID-VkImageViewCreateInfo-usage-04550"
  id="VUID-VkImageViewCreateInfo-usage-04550"></a>
  VUID-VkImageViewCreateInfo-usage-04550  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is enabled, and the `usage` for the image view includes
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, then the
  image view’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-view-format-features"
  target="_blank" rel="noopener">format features</a> **must** contain
  `VK_FORMAT_FEATURE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkImageViewCreateInfo-usage-04551"
  id="VUID-VkImageViewCreateInfo-usage-04551"></a>
  VUID-VkImageViewCreateInfo-usage-04551  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is enabled, the `usage` for the image view includes
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-layeredShadingRateAttachments"
  target="_blank"
  rel="noopener"><code>layeredShadingRateAttachments</code></a> is
  `VK_FALSE`, `subresourceRange.layerCount` **must** be `1`

- <a href="#VUID-VkImageViewCreateInfo-flags-02572"
  id="VUID-VkImageViewCreateInfo-flags-02572"></a>
  VUID-VkImageViewCreateInfo-flags-02572  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapDynamic"
  target="_blank"
  rel="noopener"><code>fragmentDensityMapDynamic</code></a> feature is
  not enabled, `flags` **must** not contain
  `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT`

- <a href="#VUID-VkImageViewCreateInfo-flags-03567"
  id="VUID-VkImageViewCreateInfo-flags-03567"></a>
  VUID-VkImageViewCreateInfo-flags-03567  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapDeferred"
  target="_blank"
  rel="noopener"><code>fragmentDensityMapDeferred</code></a> feature is
  not enabled, `flags` **must** not contain
  `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT`

- <a href="#VUID-VkImageViewCreateInfo-flags-03568"
  id="VUID-VkImageViewCreateInfo-flags-03568"></a>
  VUID-VkImageViewCreateInfo-flags-03568  
  If `flags` contains
  `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT`, `flags`
  **must** not contain
  `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT`

- <a href="#VUID-VkImageViewCreateInfo-image-03569"
  id="VUID-VkImageViewCreateInfo-image-03569"></a>
  VUID-VkImageViewCreateInfo-image-03569  
  If `image` was created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT` and `usage` containing
  `VK_IMAGE_USAGE_SAMPLED_BIT`, `subresourceRange.layerCount` **must**
  be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubsampledArrayLayers"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceFragmentDensityMap2PropertiesEXT</code>::<code>maxSubsampledArrayLayers</code></a>

- <a href="#VUID-VkImageViewCreateInfo-invocationMask-04993"
  id="VUID-VkImageViewCreateInfo-invocationMask-04993"></a>
  VUID-VkImageViewCreateInfo-invocationMask-04993  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-invocationMask"
  target="_blank" rel="noopener"><code>invocationMask</code></a> feature
  is enabled, and if `image` was created with `usage` containing
  `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI`, `format` **must** be
  `VK_FORMAT_R8_UINT`

- <a href="#VUID-VkImageViewCreateInfo-flags-04116"
  id="VUID-VkImageViewCreateInfo-flags-04116"></a>
  VUID-VkImageViewCreateInfo-flags-04116  
  If `flags` does not contain
  `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` and
  `image` was created with `usage` containing
  `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`, its `flags` **must**
  not contain any of `VK_IMAGE_CREATE_PROTECTED_BIT`,
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`,
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`

- <a href="#VUID-VkImageViewCreateInfo-pNext-02662"
  id="VUID-VkImageViewCreateInfo-pNext-02662"></a>
  VUID-VkImageViewCreateInfo-pNext-02662  
  If the `pNext` chain includes a
  [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html)
  structure, and `image` was not created with a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure included in the `pNext` chain of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), its `usage` member
  **must** not include any bits that were not set in the `usage` member
  of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure used to
  create `image`

- <a href="#VUID-VkImageViewCreateInfo-pNext-02663"
  id="VUID-VkImageViewCreateInfo-pNext-02663"></a>
  VUID-VkImageViewCreateInfo-pNext-02663  
  If the `pNext` chain includes a
  [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html)
  structure, `image` was created with a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure included in the `pNext` chain of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), and
  `subresourceRange.aspectMask` includes `VK_IMAGE_ASPECT_STENCIL_BIT`,
  the `usage` member of the
  [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html)
  structure **must** not include any bits that were not set in the
  `usage` member of the
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure used to create `image`

- <a href="#VUID-VkImageViewCreateInfo-pNext-02664"
  id="VUID-VkImageViewCreateInfo-pNext-02664"></a>
  VUID-VkImageViewCreateInfo-pNext-02664  
  If the `pNext` chain includes a
  [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html)
  structure, `image` was created with a
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)
  structure included in the `pNext` chain of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), and
  `subresourceRange.aspectMask` includes bits other than
  `VK_IMAGE_ASPECT_STENCIL_BIT`, the `usage` member of the
  [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html)
  structure **must** not include any bits that were not set in the
  `usage` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  structure used to create `image`

- <a href="#VUID-VkImageViewCreateInfo-imageViewType-04973"
  id="VUID-VkImageViewCreateInfo-imageViewType-04973"></a>
  VUID-VkImageViewCreateInfo-imageViewType-04973  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_1D`, `VK_IMAGE_VIEW_TYPE_2D`, or
  `VK_IMAGE_VIEW_TYPE_3D`; and `subresourceRange.layerCount` is not
  `VK_REMAINING_ARRAY_LAYERS`, then `subresourceRange.layerCount`
  **must** be 1

- <a href="#VUID-VkImageViewCreateInfo-imageViewType-04974"
  id="VUID-VkImageViewCreateInfo-imageViewType-04974"></a>
  VUID-VkImageViewCreateInfo-imageViewType-04974  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_1D`, `VK_IMAGE_VIEW_TYPE_2D`, or
  `VK_IMAGE_VIEW_TYPE_3D`; and `subresourceRange.layerCount` is
  `VK_REMAINING_ARRAY_LAYERS`, then the remaining number of layers
  **must** be 1

- <a href="#VUID-VkImageViewCreateInfo-viewType-02960"
  id="VUID-VkImageViewCreateInfo-viewType-02960"></a>
  VUID-VkImageViewCreateInfo-viewType-02960  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE` and
  `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `subresourceRange.layerCount` **must** be `6`

- <a href="#VUID-VkImageViewCreateInfo-viewType-02961"
  id="VUID-VkImageViewCreateInfo-viewType-02961"></a>
  VUID-VkImageViewCreateInfo-viewType-02961  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY` and
  `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `subresourceRange.layerCount` **must** be a multiple of `6`

- <a href="#VUID-VkImageViewCreateInfo-viewType-02962"
  id="VUID-VkImageViewCreateInfo-viewType-02962"></a>
  VUID-VkImageViewCreateInfo-viewType-02962  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE` and
  `subresourceRange.layerCount` is `VK_REMAINING_ARRAY_LAYERS`, the
  remaining number of layers **must** be `6`

- <a href="#VUID-VkImageViewCreateInfo-viewType-02963"
  id="VUID-VkImageViewCreateInfo-viewType-02963"></a>
  VUID-VkImageViewCreateInfo-viewType-02963  
  If `viewType` is `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY` and
  `subresourceRange.layerCount` is `VK_REMAINING_ARRAY_LAYERS`, the
  remaining number of layers **must** be a multiple of `6`

- <a href="#VUID-VkImageViewCreateInfo-imageViewFormatSwizzle-04465"
  id="VUID-VkImageViewCreateInfo-imageViewFormatSwizzle-04465"></a>
  VUID-VkImageViewCreateInfo-imageViewFormatSwizzle-04465  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`imageViewFormatSwizzle`
  is `VK_FALSE`, all elements of `components` **must** have the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a
  href="#VUID-VkImageViewCreateInfo-imageViewFormatReinterpretation-04466"
  id="VUID-VkImageViewCreateInfo-imageViewFormatReinterpretation-04466"></a>
  VUID-VkImageViewCreateInfo-imageViewFormatReinterpretation-04466  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`imageViewFormatReinterpretation`
  is `VK_FALSE`, the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) in `format` **must** not
  contain a different number of components, or a different number of
  bits in each component, than the format of the `VkImage` in `image`

- <a href="#VUID-VkImageViewCreateInfo-image-04817"
  id="VUID-VkImageViewCreateInfo-image-04817"></a>
  VUID-VkImageViewCreateInfo-image-04817  
  If `image` was created with `usage` containing
  `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`,
  `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`, or
  `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, then the `viewType`
  **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-image-04818"
  id="VUID-VkImageViewCreateInfo-image-04818"></a>
  VUID-VkImageViewCreateInfo-image-04818  
  If `image` was created with `usage` containing
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`,
  `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`, or
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`, then the `viewType`
  **must** be `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-flags-08106"
  id="VUID-VkImageViewCreateInfo-flags-08106"></a>
  VUID-VkImageViewCreateInfo-flags-08106  
  If `flags` includes
  `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBufferCaptureReplay"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-VkImageViewCreateInfo-pNext-08107"
  id="VUID-VkImageViewCreateInfo-pNext-08107"></a>
  VUID-VkImageViewCreateInfo-pNext-08107  
  If the `pNext` chain includes a
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  structure, `flags` **must** contain
  `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06787"
  id="VUID-VkImageViewCreateInfo-pNext-06787"></a>
  VUID-VkImageViewCreateInfo-pNext-06787  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06944"
  id="VUID-VkImageViewCreateInfo-pNext-06944"></a>
  VUID-VkImageViewCreateInfo-pNext-06944  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-textureSampleWeighted"
  target="_blank" rel="noopener"><code>textureSampleWeighted</code></a>
  feature **must** be enabled

- <a href="#VUID-VkImageViewCreateInfo-pNext-06945"
  id="VUID-VkImageViewCreateInfo-pNext-06945"></a>
  VUID-VkImageViewCreateInfo-pNext-06945  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure, then `image` **must** have been created with `usage`
  containing `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06946"
  id="VUID-VkImageViewCreateInfo-pNext-06946"></a>
  VUID-VkImageViewCreateInfo-pNext-06946  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure, then `components` **must** be
  `VK_COMPONENT_SWIZZLE_IDENTITY` for all components

- <a href="#VUID-VkImageViewCreateInfo-pNext-06947"
  id="VUID-VkImageViewCreateInfo-pNext-06947"></a>
  VUID-VkImageViewCreateInfo-pNext-06947  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure, then `subresourceRange.aspectMask` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06948"
  id="VUID-VkImageViewCreateInfo-pNext-06948"></a>
  VUID-VkImageViewCreateInfo-pNext-06948  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure, then `subresourceRange.levelCount` **must** be `1`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06949"
  id="VUID-VkImageViewCreateInfo-pNext-06949"></a>
  VUID-VkImageViewCreateInfo-pNext-06949  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure, then `viewType` **must** be `VK_IMAGE_VIEW_TYPE_1D_ARRAY`
  or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06950"
  id="VUID-VkImageViewCreateInfo-pNext-06950"></a>
  VUID-VkImageViewCreateInfo-pNext-06950  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and if `viewType` is `VK_IMAGE_VIEW_TYPE_1D_ARRAY`, then
  `image` **must** have been created with `imageType` `VK_IMAGE_TYPE_1D`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06951"
  id="VUID-VkImageViewCreateInfo-pNext-06951"></a>
  VUID-VkImageViewCreateInfo-pNext-06951  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and `viewType` is `VK_IMAGE_VIEW_TYPE_1D_ARRAY`, then
  `subresourceRange.layerCount` **must** be equal to `2`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06952"
  id="VUID-VkImageViewCreateInfo-pNext-06952"></a>
  VUID-VkImageViewCreateInfo-pNext-06952  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and `viewType` is `VK_IMAGE_VIEW_TYPE_1D_ARRAY`, then
  `image` **must** have been created with `width` equal to or greater
  than (numPhases×max(align(filterSize.width,4),filterSize.height))

- <a href="#VUID-VkImageViewCreateInfo-pNext-06953"
  id="VUID-VkImageViewCreateInfo-pNext-06953"></a>
  VUID-VkImageViewCreateInfo-pNext-06953  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and if `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then
  `image` **must** have been created with `imageType` `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06954"
  id="VUID-VkImageViewCreateInfo-pNext-06954"></a>
  VUID-VkImageViewCreateInfo-pNext-06954  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then
  `subresourceRange.layerCount` **must** be equal or greater than
  numPhases

- <a href="#VUID-VkImageViewCreateInfo-pNext-06955"
  id="VUID-VkImageViewCreateInfo-pNext-06955"></a>
  VUID-VkImageViewCreateInfo-pNext-06955  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then
  `image` **must** have been created with `width` equal to or greater
  than `filterSize.width`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06956"
  id="VUID-VkImageViewCreateInfo-pNext-06956"></a>
  VUID-VkImageViewCreateInfo-pNext-06956  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure and `viewType` is `VK_IMAGE_VIEW_TYPE_2D_ARRAY`, then
  `image` **must** have been created with `height` equal to or greater
  than `filterSize.height`

- <a href="#VUID-VkImageViewCreateInfo-pNext-06957"
  id="VUID-VkImageViewCreateInfo-pNext-06957"></a>
  VUID-VkImageViewCreateInfo-pNext-06957  
  If the `pNext` chain includes
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
  structure then
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)::`filterSize.height`
  **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-weightfilter-maxdimension"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceImageProcessingPropertiesQCOM</code>::<code>maxWeightFilterDimension.height</code></a>

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-09594"
  id="VUID-VkImageViewCreateInfo-subresourceRange-09594"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-09594  
  `subresourceRange.aspectMask` **must** be valid for the `format` the
  `image` was created with

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewCreateInfo-sType-sType"
  id="VUID-VkImageViewCreateInfo-sType-sType"></a>
  VUID-VkImageViewCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO`

- <a href="#VUID-VkImageViewCreateInfo-pNext-pNext"
  id="VUID-VkImageViewCreateInfo-pNext-pNext"></a>
  VUID-VkImageViewCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
  [VkImageViewASTCDecodeModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewASTCDecodeModeEXT.html),
  [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewMinLodCreateInfoEXT.html),
  [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html),
  [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSlicedCreateInfoEXT.html),
  [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html),
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html),
  or [VkSamplerYcbcrConversionInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html)

- <a href="#VUID-VkImageViewCreateInfo-sType-unique"
  id="VUID-VkImageViewCreateInfo-sType-unique"></a>
  VUID-VkImageViewCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

- <a href="#VUID-VkImageViewCreateInfo-flags-parameter"
  id="VUID-VkImageViewCreateInfo-flags-parameter"></a>
  VUID-VkImageViewCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlagBits.html) values

- <a href="#VUID-VkImageViewCreateInfo-image-parameter"
  id="VUID-VkImageViewCreateInfo-image-parameter"></a>
  VUID-VkImageViewCreateInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkImageViewCreateInfo-viewType-parameter"
  id="VUID-VkImageViewCreateInfo-viewType-parameter"></a>
  VUID-VkImageViewCreateInfo-viewType-parameter  
  `viewType` **must** be a valid [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html)
  value

- <a href="#VUID-VkImageViewCreateInfo-format-parameter"
  id="VUID-VkImageViewCreateInfo-format-parameter"></a>
  VUID-VkImageViewCreateInfo-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkImageViewCreateInfo-components-parameter"
  id="VUID-VkImageViewCreateInfo-components-parameter"></a>
  VUID-VkImageViewCreateInfo-components-parameter  
  `components` **must** be a valid
  [VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html) structure

- <a href="#VUID-VkImageViewCreateInfo-subresourceRange-parameter"
  id="VUID-VkImageViewCreateInfo-subresourceRange-parameter"></a>
  VUID-VkImageViewCreateInfo-subresourceRange-parameter  
  `subresourceRange` **must** be a valid
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html),
[VkImageViewCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlags.html),
[VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
